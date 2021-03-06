CREATE PROCEDURE [dbo].[InsertInterest]
      @InterestName nvarchar(50)
      ,@ForumId int
AS
BEGIN
	SET NOCOUNT ON;
	IF EXISTS (SELECT TOP 1* FROM [Theme] WHERE ThemeName = @InterestName) OR
	NOT EXISTS (SELECT * FROM [Forum] WHERE ForumId = @ForumId)
BEGIN
  RAISERROR ( 'Interest already exists or forum does not exist.',1,1)
END
ELSE
	BEGIN
		INSERT INTO [Theme] (ThemeName) 
	  SELECT  @InterestName

	 DECLARE  @InterestId int =  SCOPE_IDENTITY()

	INSERT INTO [dbo].[ThemeForum] ([ForumID], [ThemeId])
	SELECT @ForumId,  @InterestId
END
	

END

