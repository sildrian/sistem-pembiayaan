-- Create database
CREATE DATABASE testbtpns;
GO

-- Switch to the database
USE testbtpns;
GO

-- Create schema
CREATE SCHEMA tbtpns;
GO

CREATE TABLE tbtpns.[user] (
    user_id INT IDENTITY(1,1) PRIMARY KEY,
    name NVARCHAR(100) DEFAULT '',
    phone NVARCHAR(20) DEFAULT ''
);
GO

CREATE TABLE tbtpns.facility_limit (
    facility_limit_id INT IDENTITY(1,1) PRIMARY KEY,
    user_id INT DEFAULT 0,
    limit_amount INT DEFAULT 0,
    CONSTRAINT FK_facility_limit_user FOREIGN KEY (user_id)
        REFERENCES tbtpns.[user](user_id)
);
GO

CREATE TABLE tbtpns.tenor (
    tenor_id INT IDENTITY(1,1) PRIMARY KEY,
    tenor_value NVARCHAR(20) DEFAULT '6,12,18,24,30,36'
);
GO

CREATE TABLE tbtpns.user_facility (
    user_facility_id INT IDENTITY(1,1) PRIMARY KEY,
    user_id INT DEFAULT 0,
    facility_limit_id INT DEFAULT 0,
    amount INT DEFAULT 0,
    tenor INT DEFAULT 0,
    start_date NVARCHAR(10) DEFAULT '',
    monthly_installment INT DEFAULT 0,
    total_margin INT DEFAULT 0,
    total_payment INT DEFAULT 0,
    CONSTRAINT FK_user_facility_user FOREIGN KEY (user_id)
        REFERENCES tbtpns.[user](user_id),
    CONSTRAINT FK_user_facility_facility_limit FOREIGN KEY (facility_limit_id)
        REFERENCES tbtpns.facility_limit(facility_limit_id)
);
GO

CREATE TABLE tbtpns.user_facility_details (
    detail_id INT IDENTITY(1,1) PRIMARY KEY,
    user_facility_id INT DEFAULT 0,
    due_date NVARCHAR(10) DEFAULT '',
    installment_amount INT DEFAULT 0,
    CONSTRAINT FK_user_facility_details_user_facility FOREIGN KEY (user_facility_id)
        REFERENCES tbtpns.user_facility(user_facility_id)
);
GO
