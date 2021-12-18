@echo off

call run :init_colors

:check_run_mode
    if "%*"=="" (       
        call :run_executable 
    ) else (
        call :run_library %*
    )
    goto :eof

:run_executable
    rem run stuff.bat as executable script
    call run :color ‘%CL_INFO%‘ ‘This is not executable script. Please use 'run.bat' only.‘
    call run :ask   
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
    call run :color ‘%CL_INFO%‘ ‘PATH=%PATH%‘
    call run :color ‘%CL_INFO%‘ ‘GOPATH=%GOPATH%‘

    set ARCH_URL=https://golang.org/dl/go1.16.5.windows-amd64.zip
    set ARCH_FOLDER=go
    goto :eof

:install
    call run :install golang %ARCH_URL% %ARCH_FOLDER%
    goto :eof

:version
    call run :eval_echo_color ‘%GO% version‘
    goto :eof

:build
    rem do nothing
    goto :eof

:test    
    call run :eval_echo ‘%GO% test ./...‘
    echo.
    goto :eof

:run
    call run :eval_echo ‘%GO% run main.go %GAME_TO_RUN% %SERVER_URL%‘
    goto :eof