set +x

nvim \
-c ":nmap <M-r> :!go run %:p<CR>" $@
