import express, { Application, Request, Response } from "express";
import sequelize from "sequelize";
import { Sequelize, DataTypes, QueryInterface } from "sequelize";
const app: Application = express();
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// 環境変数
const mysqlDatabase = process.env.MYSQL_DATABASE; // データベース名
const mysqlRootPassword = process.env.MYSQL_ROOT_PASSWORD; // ROOTパスワード
const dbHost = process.env.DB_CONTAINER_IPV4; // DBコンテナIPv4
const dbPort = process.env.DB_CONTAINER_PORT; // DBコンテナポート番号
const apiHost = process.env.API_CONTAINER_IPV4; // APIコンテナIPv4
const apiPort = process.env.API_CONTAINER_PORT; // APIコンテナポート番号
const modelName = "User"; // モデル名

// DBに接続する関数
function accessDB() {
  if (
    mysqlDatabase === undefined ||
    mysqlRootPassword === undefined ||
    dbHost === undefined ||
    dbPort === undefined
  ) {
    return null;
  }
  const sequelize = new Sequelize(mysqlDatabase, "root", mysqlRootPassword, {
    host: dbHost,
    port: parseInt(dbPort),
    dialect: "mysql",
    logging: false,
  });
  return sequelize;
}

// テーブルの定義関数
function defineModel(sequelize: Sequelize) {
  const model = sequelize.define(modelName, {
    // Userテーブル
    id: {
      type: DataTypes.INTEGER,
      primaryKey: true,
      autoIncrement: true, // AUTO_INCREMENT
      allowNull: false, // Not Null
      comment: "ID, 主キー",
    },
    name: {
      type: DataTypes.STRING, // 文字列型
      allowNull: false, // Not Null
    },
    age: {
      type: DataTypes.INTEGER, // 整数型
    },
  });
  return model;
}

// テーブルの作成
app.post("/create_table", async (_req: Request, res: Response) => {
  // DB接続
  const sequelize = accessDB();
  if (sequelize === null) {
    return res.status(500).send({
      message: "[ERROR] Access DB",
    });
  }
  // テーブルの定義
  const model = defineModel(sequelize);
  try {
    // テーブルの同期
    await model.sync({
      force: false, // true:テーブルが存在する場合、削除した上で新規作成する, false:テーブルが存在する場合は何もしない
    });
  } catch (e) {
    return res.status(500).send({
      message: e,
    });
  }
  return res.status(200).send({
    message: "Success to create table.",
  });
});

// INSERT
app.put("/insert", async (req: Request, res: Response) => {
  // POSTデータ取得
  const body = req.body;
  if (!("name" in body) || body.name === "") {
    return res.status(500).send({
      message: "[ERROR] No name",
    });
  }
  if (!("age" in body) || body.age === "") {
    return res.status(500).send({
      message: "[ERROR] No age",
    });
  }
  console.log(body);
  // DB接続
  const sequelize = accessDB();
  if (sequelize === null) {
    return res.status(500).send({
      message: "[ERROR] Access DB",
    });
  }
  // テーブルの定義
  const model = defineModel(sequelize);
  try {
    // テーブルの同期
    await model.sync({
      force: false, // true:テーブルが存在する場合、削除した上で新規作成する, false:テーブルが存在する場合は何もしない
    });
  } catch (e) {
    return res.status(500).send({
      message: e,
    });
  }
  try {
    // INSERTの実行
    const user = await model.create({
      name: body.name,
      age: body.age,
    });
  } catch (e) {
    return res.status(500).send({
      message: e,
    });
  }
  return res.status(200).send({
    message: "Success to insert data.",
  });
});

// SELECT
app.get("/select", async (_req: Request, res: Response) => {
  const sequelize = accessDB();
  try {
    // 生クエリの実行
    const rows = await sequelize?.query("SELECT * FROM Users");
    return res.status(200).send({
      message: rows,
    });
  } catch (e) {
    return res.status(500).send({
      message: e,
    });
  }
});

// テーブル削除
app.delete("/drop_table", async (_req: Request, res: Response) => {
  // DB接続
  const sequelize = accessDB();
  if (sequelize === null) {
    return res.status(500).send({
      message: "[ERROR] Access DB",
    });
  }
  // テーブルの定義
  const model = defineModel(sequelize);
  try {
    // テーブルの同期
    await model.sync({
      force: false, // true:テーブルが存在する場合、削除した上で新規作成する, false:テーブルが存在する場合は何もしない
    });
  } catch (e) {
    return res.status(500).send({
      message: e,
    });
  }
  try {
    // テーブルの削除
    await model.drop();
  } catch (e) {
    return res.status(500).send({
      message: e,
    });
  }
  return res.status(200).send({
    message: "Success to delete table.",
  });
});

// サーバーを起動する処理
try {
  app.listen(apiPort, () => {
    console.log("server running at: http://" + apiHost + ":" + apiPort);
  });
} catch (e) {
  if (e instanceof Error) {
    console.error(e.message);
  }
}
