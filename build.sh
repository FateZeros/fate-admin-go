#!/bin/bash
#=============================================================================
# Author: yejianfei
#
# Filename:		build.sh
#=============================================================================


function usage {
    cat << EOF
  echo -e "\nUsage: $0 (install|start|stop)"    
    echo "Examle:     
    bash $0 install 
EOF
}

function prepare_check {
    cat << EOF

      执行此脚本之前，请确认一下软件是否安装或者是否有现成的连接地址。

      若未没有请根据不同的系统，自行百度一下安装教程。

          1. git 最新版本即可
          1. MySQL >= 5.7
          2. Go >= 1.14
          3. Redis 最新版本即可
          4. node >= v12 （稳定版本）
          5. npm >= v6.14.8
          6. 安装docker

EOF
    check_soft docker
    check_soft git
    check_soft go
    check_soft npm
}

function check_soft {
    local _soft_name=$1
    command -v ${_soft_name} > /dev/null || {
    echo_red "请安装 ${_soft_name} 后再执行本脚本安装ferry。"
        exit 1
    }
}

function install_app() {
    prepare_check
    # init
    # get_variables
    # install_front
    # install_backend
}

function main {
    action=${1-}
    target=${2-}
    args=("$@")

    case "${action}" in
        install)
            install_app
            ;;
        *)
            echo "No such command: ${action}"
            usage
            ;;
    esac
}

main "$@"