set +x

nvim \
-c ":nmap <M-r> :!cargo run<CR>" $@
