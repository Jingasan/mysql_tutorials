package main

// Gin/GORMのimport
import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// レコードの構造体定義
type User struct {
	// レコードにid, created_at, updatede_at, deleted_atが自動で追加される
	gorm.Model
	Name string `label:"ユーザー名" gorm:"column:name;type:varchar(64)" json:"name"`
	Age  uint   `label:"年齢" gorm:"column:age;type:int" json:"age"`
}
type Organization struct {
	gorm.Model
	Name string `label:"部署名" gorm:"column:name;type:varchar(64)" json:"name"`
}

// DBとの接続を確立する関数
func connectDB() (*gorm.DB, error) {
	database_name := os.Getenv("MYSQL_DATABASE")        // データベース名
	root_password := os.Getenv("MYSQL_ROOT_PASSWORD")   // ROOTパスワード
	db_container_ipv4 := os.Getenv("DB_CONTAINER_IPV4") // DBコンテナIPv4
	db_container_port := os.Getenv("DB_CONTAINER_PORT") // DBコンテナポート番号
	dsn := "root:" + root_password + "@tcp(" + db_container_ipv4 + ":" + db_container_port + ")/" + database_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func main() {
	// GinをReleaseモードに設定
	gin.SetMode(gin.ReleaseMode)

	// Engineインスタンスの作成
	engine := gin.Default()

	// テーブルの作成メソッド
	engine.POST("/table", func(ctx *gin.Context) {
		// DBとの接続を確立
		db, err := connectDB()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		// 処理完了後にDBとの接続をクローズ
		sqlDb, err := db.DB()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		defer sqlDb.Close()
		// テーブルの作成
		db.Migrator().CreateTable(&User{})
		db.Migrator().CreateTable(&Organization{})
		// レスポンス
		ctx.JSON(http.StatusOK, "OK")
	})

	// テーブルの削除メソッド
	engine.DELETE("/table", func(ctx *gin.Context) {
		// DBとの接続を確立
		db, err := connectDB()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		// 処理完了後にDBとの接続をクローズ
		sqlDb, err := db.DB()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		defer sqlDb.Close()
		// テーブルの削除
		db.Migrator().DropTable(&User{})
		db.Migrator().DropTable(&Organization{})
		// レスポンス
		ctx.JSON(http.StatusOK, "OK")
	})

	// 全レコードの取得メソッド
	engine.GET("/user", func(ctx *gin.Context) {
		// DBとの接続を確立
		db, err := connectDB()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		// 処理完了後にDBとの接続をクローズ
		sqlDb, err := db.DB()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		defer sqlDb.Close()
		// レコードの取得
		users := []*User{}
		err = db.Find(&users).Error
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		// デバッグプリント
		for _, r := range users {
			log.Printf("id:%d, created_at:%s, updated_at:%s, deleted_at:%s, name:%s, age:%d",
				r.ID, r.CreatedAt, r.UpdatedAt, r.DeletedAt.Time, r.Name, r.Age)
		}
		// レスポンス
		ctx.JSON(http.StatusOK, users)
	})

	// レコードの挿入メソッド
	engine.POST("/user", func(ctx *gin.Context) {
		// DBとの接続を確立
		db, err := connectDB()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		// 処理完了後にDBとの接続をクローズ
		sqlDb, err := db.DB()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		defer sqlDb.Close()
		// リクエストボディのデータ型
		type RequestBodyDataType struct {
			Name string `json:"name"`
			Age  uint   `json:"Age"`
		}
		var reqJSON RequestBodyDataType
		// リクエストボディが規定の型を満たさない場合のエラー処理
		if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
			ctx.JSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			return
		}
		// レコードの挿入
		db.Create(&User{Name: reqJSON.Name, Age: reqJSON.Age})
		// レスポンス
		ctx.JSON(http.StatusOK, "OK")
	})

	// 指定レコードの取得メソッド
	engine.GET("/user/:id", func(ctx *gin.Context) {
		// DBとの接続を確立
		db, err := connectDB()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		// 処理完了後にDBとの接続をクローズ
		sqlDb, err := db.DB()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		defer sqlDb.Close()
		// URLパラメータからレコード取得IDを取得
		id := ctx.Param("id")
		// レコードの取得
		user := User{}
		err = db.First(&user, id).Error
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		// レスポンス
		ctx.JSON(http.StatusOK, user)
	})

	// 指定レコードの更新メソッド
	engine.PUT("/user/:id", func(ctx *gin.Context) {
		// DBとの接続を確立
		db, err := connectDB()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		// 処理完了後にDBとの接続をクローズ
		sqlDb, err := db.DB()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		defer sqlDb.Close()
		// URLパラメータからレコード取得IDを取得
		id := ctx.Param("id")
		// リクエストボディのデータ型
		type RequestBodyDataType struct {
			Name string `json:"name"`
			Age  uint   `json:"Age"`
		}
		var reqJSON RequestBodyDataType
		// リクエストボディが規定の型を満たさない場合のエラー処理
		if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
			ctx.JSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			return
		}
		// レコードの更新
		user := User{}
		err = db.Model(&user).Where("id = ?", id).Updates(User{Name: reqJSON.Name, Age: reqJSON.Age}).Error
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		// レスポンス
		ctx.JSON(http.StatusOK, "OK")
	})

	// レコードの削除
	engine.DELETE("/user/:id", func(ctx *gin.Context) {
		// DBとの接続を確立
		db, err := connectDB()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		// 処理完了後にDBとの接続をクローズ
		sqlDb, err := db.DB()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		defer sqlDb.Close()
		// URLパラメータから削除するIDを取得
		id := ctx.Param("id")
		// レコードの削除
		err = db.Where("id = ?", id).Delete(&User{}).Error
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		// レスポンス
		ctx.JSON(http.StatusOK, "OK")
	})

	// 指定レコードの指定カラムの取得メソッド
	engine.GET("/user/:id/:column", func(ctx *gin.Context) {
		// DBとの接続を確立
		db, err := connectDB()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		// 処理完了後にDBとの接続をクローズ
		sqlDb, err := db.DB()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		defer sqlDb.Close()
		// URLパラメータからレコード取得IDを取得
		id := ctx.Param("id")
		column := ctx.Param("column")
		// レコードの取得
		user := User{}
		db.Raw("SELECT " + column + " FROM users WHERE id = " + id).Scan(&user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		// レスポンス
		if column == "name" {
			ctx.JSON(http.StatusOK, gin.H{"name": user.Name})
		} else if column == "age" {
			ctx.JSON(http.StatusOK, gin.H{"age": user.Age})
		} else {
			ctx.JSON(http.StatusBadRequest,
				gin.H{"error": "Bad URL parameter"})
		}
	})

	// ポート番号を指定
	apiHost := os.Getenv("API_CONTAINER_IPV4") // APIコンテナIPv4
	apiPort := os.Getenv("API_CONTAINER_PORT") // APIコンテナポート番号
	engine.Run(":" + apiPort)
	log.Print("server running at: http://" + apiHost + ":" + apiPort)
}
