USE [Forum]
GO
/****** Object:  StoredProcedure [dbo].[InsertUser]    Script Date: 20.11.2021 18:29:52 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
ALTER PROCEDURE [dbo].[InsertUser]
	  @UserType smallint
      ,@UserName nvarchar(50)
      ,@UserMail nvarchar(50)
      ,@Password nvarchar(50)
	  ,@Interests NTEXT
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;
INSERT INTO [dbo].[User] ([UserType], [UserName], [UserMail], [Password], [IsDeleted])
  SELECT  @UserType, @UserName, @UserMail, @Password, 0

 DECLARE @IdUser int =  SCOPE_IDENTITY()

  create table #dataitems 
([Interest] nvarchar(50))

  DECLARE @xml int
 if datalength(@Interests) = 0
  exec master..sp_xml_preparedocument @xml OUTPUT, null
 else
  exec master..sp_xml_preparedocument @xml OUTPUT, @Interests

  
 INSERT INTO #dataitems([Interest])
  SELECT [Interest]
    FROM OPENXML(@xml, 'ITEMS/ITEM', 1)
  WITH(
  Interest nvarchar(50) '@Interest' 
)

INSERT INTO [dbo].[ForumUser] ([ForumID], [UserId])
SELECT dbo.ThemeForum.ForumID, @IdUser
FROM     dbo.Forum INNER JOIN
                  dbo.ThemeForum ON dbo.Forum.ForumID = dbo.ThemeForum.ForumID INNER JOIN
                  dbo.Theme ON dbo.ThemeForum.ThemeID = dbo.Theme.ThemeID
				  WHERE dbo.Theme.ThemeName IN (SELECT * FROM #dataitems) 
END

