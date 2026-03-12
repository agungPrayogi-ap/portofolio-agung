@echo off
echo Syncing views and public to api folder...
xcopy /E /I /Y "views" "api\views"
xcopy /E /I /Y "public" "api\public"
echo.
echo Sync complete! Now run:
echo   git add .
echo   git commit -m "your message"
echo   git push origin main
