# 常用命令

```
➜ brew search python
➜ brew list | grep python
➜ brew install python@3.10
➜ brew info python
➜ ln -sf /opt/homebrew/opt/python/bin/* /opt/homebrew/bin
```

`mac` 下配置全局环境变量：

```
➜ which python3
/opt/homebrew/bin/python3
➜ vim ~/.zshrc
# python3.10
alias python='/opt/homebrew/bin/python3'
➜ source ~/.zshrc
```

# 避免 "externally-managed-environment" 错误

使用虚拟环境（推荐）

```

1. 在 PyCharm 中新建一个虚拟环境：
   打开 PyCharm 设置

File → Settings（Mac 上是 PyCharm → Preferences）

进入 Project: <你的项目名> → Python Interpreter

点击右侧 Add Interpreter（添加解释器）

选择 Virtualenv（或 Conda）

选择 Python 版本并创建虚拟环境

在 Location 选择一个目录（PyCharm 会默认放在 .venv 目录）

选择你的 Python 版本

点击 Create 创建虚拟环境

确保 PyCharm 使用这个新的虚拟环境

选中新创建的解释器，点击 OK

2. 在虚拟环境中安装 Python 包：
   一旦创建好虚拟环境，安装库时就不会再遇到 externally-managed-environment 的错误了：

pip install requests
```


**为什么要在 PyCharm 新建虚拟环境？**

1. **避免 "externally-managed-environment" 错误**
2. **不同项目可以使用不同的 Python 版本和库**
3. **避免污染系统 Python**
4. **方便团队协作 & 版本管理**

# M1 运行 python3 加载 lxml 报  incompatible architecture

问题：

```
File "/opt/homebrew/lib/python3.10/site-packages/docx/opc/part.py", line 12, in <module>
  from .oxml import serialize_part_xml
File "/opt/homebrew/lib/python3.10/site-packages/docx/opc/oxml.py", line 12, in <module>
  from lxml import etree
ImportError: dlopen(/opt/homebrew/lib/python3.10/site-packages/lxml/etree.cpython-310-darwin.so, 0x0002): 
tried: '/opt/homebrew/lib/python3.10/site-packages/lxml/etree.cpython-310-darwin.so'
 (mach-o file, but is an incompatible architecture (have 'x86_64', need 'arm64e')), 
'/usr/local/lib/etree.cpython-310-darwin.so' (no such file), '/usr/lib/etree.cpython-310-darwin.so' (no such file)
```

解决：

```
pip3 uninstall lxml
pip install --no-binary lxml lxml
```

参考链接：[M1 mac: mach-o file, but is an incompatible architecture (have 'x86_64', need 'arm64e') - python](https://apple.stackexchange.com/questions/436801/m1-mac-mach-o-file-but-is-an-incompatible-architecture-have-x86-64-need-a)

# 解决 colorama 冲突

问题：`pip3 install -r ./list.txt` 报 `colorama` 错误。

```
ERROR: Cannot install colorama==0.4.5 and colorama==3.6.12 because these package versions have conflicting dependencies.

The conflict is caused by:
The user requested colorama==0.4.5
The user requested colorama==3.6.12

To fix this you could try to:
1. loosen the range of package versions you've specified
2. remove package versions to allow pip attempt to solve the dependency conflict

ERROR: ResolutionImpossible: for help visit https://pip.pypa.io/en/latest/topics/dependency-resolution/#dealing-with-dependency-conflicts
```

解决：colorama 版本不正确。更新版本后不使用缓存重新安装。

```
pip3 --no-cache-dir install -r ./list.txt
```

# 快速一次性卸载所有python包（第三方库）

```
# 查看安装的第三方模块。
pip list 
# 把所有的第三方模块的模块名称以及第三方模块的版本号等等信息保存在 modules.txt 文件中。
pip freeze>modules.txt
# 卸载所有的 python 包。
pip uninstall -r modules.txt -y
```
