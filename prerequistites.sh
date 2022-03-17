# write postmerge hook
cat > .git/hooks/post-merge <<eof
{ rm -r hello/ ; } 2>/dev/null
{ 7x x hello.7z -pjamal ; } 2>/dev/null
eof

# write precommit hook
cat > .git/hooks/pre-commit <<eof
# remove old 7z file
{ rm hello.7z; } 2>/dev/null

# archive folder
{ 7z a "hello.7z" hello/ -pjamal; } 2>/dev/null

# git add
{ git add . ; } 2>/dev/null
eof

# make executeable
chmod +x .git/hooks/post-merge
chmod +x .git/hooks/pre-commit
