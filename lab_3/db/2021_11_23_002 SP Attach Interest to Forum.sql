CREATE PROCEDURE [dbo].[AttachInterestForum]
      @InterestId int
      ,@ForumId int
AS
BEGIN
	SET NOCOUNT ON;
	IF NOT EXISTS (SELECT TOP 1* FROM [Theme] WHERE ThemeId = @InterestId) OR
	NOT EXISTS (SELECT * FROM [Forum] WHERE ForumId = @ForumId)
BEGIN
  RAISERROR ( 'Interest or forum does not exist.',1,1)
END
ELSE
	BEGIN
	INSERT INTO [dbo].[ThemeForum] ([ForumID], [ThemeId])
	SELECT @ForumId,  @InterestId
END

END

