
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