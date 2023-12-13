active_envs=("$HOME/.envrc")

for envrc in "${active_envs[@]}"; do
    if [ -f "$envrc" ]; then
        source "$envrc"
    fi
done