# git reset

# 修正提交
*  git commit --amend // 将暂存区中的文件提交（先可以将需要提交的文件进行add）

# 本地文件撤销操作
* git checkout -- `<file>` // 撤销本地文件的所有修改
* git reset HEAD `<file> `// 取消暂存的文件


# git stash 

git stash 命令：临时暂存和恢复

* git stash list // 列出暂存的列表
* git stash save "备注信息"  // 执行暂存并添加备注
* git stash pop // 恢复暂存的工作目录，并将缓存堆栈中的对应stash删除，将对应修改应用到当前的工作目录下，默认为第一个stash,即stash@{0}，如果要应用并删除其他stash，命令：git stash pop stash@{$num} ，比如应用并删除第二个：git stash pop stash@{1}
* git stash show // 显示做了哪些改动，默认show第一个存储，如果要显示其他存贮，后面加stash@{$num}，比如第二个 git stash show stash@{1}
* git stash show -p  // 显示第一个存储的改动，如果想显示其他存存储，命令：git stash show  stash@{$num}  -p ，比如第二个：git stash show  stash@{1}  -p
* git stash apply  // 应用某个存储，但不会把存储从存储列表中删除，默认使用第一个存储,即stash@{0}，如果要使用其他个，git stash apply stash@{$num} ， 比如第二个：git stash apply stash@{1} 
* git stash drop stash@{$num} // 丢弃stash@{$num}存储，并从列表中删除这个存储
* git stash clear // 删除所有缓存的stash


# 查看远程分支

* git branch -r | grep x // 查看远程分支并进行正则匹配
*  git remote // 不带参数，列出已经存在的远程分支
* git remote -v // 列出详细信息，在每一个名字后面列出其远程url，此时， -v 选项，显示对应的克隆地址
* git remote show origin // 查看远程库的一些信息，及与本地分支的信息

# git 拉取远程分支并切换到该分支

* git checkout -b 本地分支名 origin/远程分支名
* git checkout --track origin/远程分支名 （这种写法是上面的简化版，效果完全一样）
* git checkout -t origin/远程分支名（这种写法是上面两种的简化版）


# git 查看所有远程分支以及同步

* git branch   // 查看本地仓库的分支
* git branch -a  // 查看本地和远程仓库的所有分支
* git branch -r  // 查看远程仓库的分支
* git fetch // 将本地分支与远程保持同步

# 删除分支
* git branch -D 本地分支名 // 删除本地分支
* git push origin --delete 远程分支名 // 删除远程分支
