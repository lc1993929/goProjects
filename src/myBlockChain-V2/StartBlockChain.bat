cd F:\liuchang\goProjects\src\myBlockChain-V2
del myBlockChain.db
del myBlockChain.db.lock
del block.exe
go build -o block.exe
block.exe printChain