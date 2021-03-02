__bashcompembed_custom_func() {
    #COMP_CWORD=1 current arg where user tabtab
    #COMP_WORDS=([0]="saysomething" [1]="hello" [2]="aaa")
    #COMP_LINE='saysomething hello aaa'

    COMMANDS="hello goodbye thanks completion"
    # Main Command
    if [ ${COMP_CWORD} -eq 1 ]; then
        COMPREPLY=($(compgen -W "${COMMANDS}" "${COMP_WORDS[1]}"))
        return
    fi

    case ${COMP_WORDS[1]} in
        thanks)
            [ ${COMP_CWORD} -ne 2 ] && return
            THX="from to"
            COMPREPLY=($(compgen -W "${THX}" "${COMP_WORDS[2]}"))
            ;;
    esac
}
