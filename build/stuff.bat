@echo off

if "%RUN%"=="" set RUN=%CD%\run
if "%STUFF%"=="" set STUFF=%CD%\stuff

call %RUN% :init_colors

:check_run_mode
    if "%*"=="" (       
        call :run_executable 
    ) else (
        call :run_library %*
    )
    goto :eof

:run_executable
    rem run stuff.bat as executable script
    call %RUN% :color ‘%CL_INFO%‘ ‘This is not executable script. Please use 'run.bat' only.‘
    call %RUN% :ask   
    goto :eof

:run_library
    rem run stuff.bat as library
    call %*     
    goto :eof          

:settings
    if "%INSTALL_LOCALLY%"=="true" ( set GOPATH=)

    if "%GOPATH%"==""     ( set NO_GO=true)
    if "%NO_GO%"=="true"  ( set GOPATH=%ROOT%\.golang)
    if "%NO_GO%"=="true"  ( set PATH=%GOPATH%\bin;%PATH%)

    set GO=%GOPATH%\bin\go

    echo Language environment variables
    call %RUN% :color ‘%CL_INFO%‘ ‘PATH=%PATH%‘
    call %RUN% :color ‘%CL_INFO%‘ ‘GOPATH=%GOPATH%‘

    set ARCH_URL=https://golang.org/dl/go1.16.5.windows-amd64.zip
    set ARCH_FOLDER=go
    goto :eof

:install
    call %RUN% :install golang %ARCH_URL% %ARCH_FOLDER%
    goto :eof

:version
    call %RUN% :eval_echo_color ‘%GO% version‘
    goto :eof

:build
    rem do nothing
    goto :eof

:test    
    call %RUN% :eval_echo ‘%GO% test %ROOT%\...‘
    echo.
    goto :eof

:run
    call %RUN% :eval_echo ‘%GO% run %ROOT%\main.go %GAME_TO_RUN% %SERVER_URL%‘
    goto :eof