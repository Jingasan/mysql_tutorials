CREATE TABLE
	departments (
		dept_id INT NOT NULL, -- 部門ID
		dept_name VARCHAR(255) NOT NULL, -- 部署名
		PRIMARY KEY (dept_id) -- 主キー
	);

CREATE TABLE
	employees (
		emp_id INT NOT NULL, -- 従業員番号
		dept_id INT NOT NULL, -- 部署番号
		emp_name VARCHAR(20) NOT NULL, -- 担当者名
		birthday DATE NOT NULL, -- 生年月日
		hiredate DATE NOT NULL, -- 入社年月日
		sex INT NOT NULL, -- 性別
		sal NUMERIC(9, 2), -- 給与額
		comm NUMERIC(7, 2), -- 歩合給
		PRIMARY KEY (emp_id), -- 主キー
		FOREIGN KEY (dept_id) -- 外部キー
		REFERENCES departments (dept_id) -- 部署テーブル.部署番号
	);