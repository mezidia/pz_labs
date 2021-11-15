CREATE PROCEDURE [dbo].[DeleteUser] 
@UserID int
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;
 UPDATE  [dbo].[User] 
 SET  [IsDeleted] = 1
 WHERE [UserID] = @UserID
END
GO
/****** Object:  StoredProcedure [dbo].[GetUserByID]    Script Date: 14.11.2021 17:18:13 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================



-- =============================================
CREATE PROCEDURE [dbo].[GetUserByID]
	 @UserID int
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;
SELECT [UserID], [UserType], [UserName], [UserMail], [IsDeleted]
  FROM [dbo].[User] WHERE UserID = @UserID AND IsDeleted=0 
END
GO
/****** Object:  StoredProcedure [dbo].[GetUsers]    Script Date: 14.11.2021 17:18:13 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================



-- =============================================
CREATE PROCEDURE [dbo].[GetUsers]
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;
SELECT [UserID], [UserType], [UserName], [UserMail], [IsDeleted]
  FROM [dbo].[User] WHERE IsDeleted=0 
END
GO
/****** Object:  StoredProcedure [dbo].[InsertUser]    Script Date: 14.11.2021 17:18:13 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================



-- =============================================
CREATE PROCEDURE [dbo].[InsertUser]
	  @UserType smallint
      ,@UserName nvarchar(50)
      ,@UserMail nvarchar(50)
      ,@Password nvarchar(50)
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;
INSERT INTO [dbo].[User] ([UserType], [UserName], [UserMail], [Password], [IsDeleted])
  SELECT  @UserType, @UserName, @UserMail, @Password, 0
END
GO
/****** Object:  StoredProcedure [dbo].[UpdateUser]    Script Date: 14.11.2021 17:18:13 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================



-- =============================================
CREATE PROCEDURE [dbo].[UpdateUser]
	   @UserID int
	  ,@UserType smallint
      ,@UserName nvarchar(50)
      ,@UserMail nvarchar(50)
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;
	UPDATE [dbo].[User]
 SET [UserType] = @UserType,
     [UserName] = @UserName,
     [UserMail] = @UserMail
  WHERE [UserID] = @UserID
END
GO
