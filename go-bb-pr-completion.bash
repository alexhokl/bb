_go_bb_pr() {
  COMPREPLY=()
  cur="${COMP_WORDS[COMP_CWORD]}"
  prev="${COMP_WORDS[COMP_CWORD-1]}"

  opts="approve checkout decline describe list merge open unapprove"

  case "${prev}" in
        *)
        ;;
    esac

   COMPREPLY=($(compgen -W "${opts}" -- ${cur}))
   return 0

}

complete -F _go_bb_pr go-bb-pr pr
