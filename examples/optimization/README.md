# Compiler Optimization
파일이 나눠지는 경우에는 inline 되지 않고 그대로 호출됨. 이 경우 LTO 필요.

## No optimization (-O0)
```sh
clang main.c -o main
objdump -d main
```
`calculation` 호출 후 반환값 저장(`mov  %eax,-0x8(%rbp)`)
```text
0000000000001140 <doubleNum>:
    1140:       55                      push   %rbp
    1141:       48 89 e5                mov    %rsp,%rbp
    1144:       89 7d fc                mov    %edi,-0x4(%rbp)
    1147:       8b 45 fc                mov    -0x4(%rbp),%eax
    114a:       c1 e0 01                shl    $0x1,%eax
    114d:       5d                      pop    %rbp
    114e:       c3                      ret    
    114f:       90                      nop

0000000000001150 <square>:
    1150:       55                      push   %rbp
    1151:       48 89 e5                mov    %rsp,%rbp
    1154:       89 7d fc                mov    %edi,-0x4(%rbp)
    1157:       8b 45 fc                mov    -0x4(%rbp),%eax
    115a:       0f af 45 fc             imul   -0x4(%rbp),%eax
    115e:       5d                      pop    %rbp
    115f:       c3                      ret    

0000000000001160 <calculation>:
    1160:       55                      push   %rbp
    1161:       48 89 e5                mov    %rsp,%rbp
    1164:       48 83 ec 10             sub    $0x10,%rsp
    1168:       89 7d fc                mov    %edi,-0x4(%rbp)
    116b:       89 75 f8                mov    %esi,-0x8(%rbp)
    116e:       8b 7d fc                mov    -0x4(%rbp),%edi
    1171:       e8 da ff ff ff          call   1150 <square>
    1176:       89 45 f4                mov    %eax,-0xc(%rbp)
    1179:       8b 7d f8                mov    -0x8(%rbp),%edi
    117c:       e8 bf ff ff ff          call   1140 <doubleNum>
    1181:       89 c1                   mov    %eax,%ecx
    1183:       8b 45 f4                mov    -0xc(%rbp),%eax
    1186:       0f af c1                imul   %ecx,%eax
    1189:       48 83 c4 10             add    $0x10,%rsp
    118d:       5d                      pop    %rbp
    118e:       c3                      ret    
    118f:       90                      nop

0000000000001190 <main>:
    1190:       55                      push   %rbp
    1191:       48 89 e5                mov    %rsp,%rbp
    1194:       48 83 ec 10             sub    $0x10,%rsp
    1198:       c7 45 fc 00 00 00 00    movl   $0x0,-0x4(%rbp)
    119f:       bf 03 00 00 00          mov    $0x3,%edi
    11a4:       be 05 00 00 00          mov    $0x5,%esi
    11a9:       e8 b2 ff ff ff          call   1160 <calculation>
    11ae:       89 45 f8                mov    %eax,-0x8(%rbp)
    11b1:       8b 75 f8                mov    -0x8(%rbp),%esi
    11b4:       48 8d 3d 49 0e 00 00    lea    0xe49(%rip),%rdi        # 2004 <_IO_stdin_used+0x4>
    11bb:       b0 00                   mov    $0x0,%al
    11bd:       e8 6e fe ff ff          call   1030 <printf@plt>
    11c2:       8b 45 f8                mov    -0x8(%rbp),%eax
    11c5:       48 83 c4 10             add    $0x10,%rsp
    11c9:       5d                      pop    %rbp
    11ca:       c3                      ret    
```

## All optimization wihtout violating strict compliance (-O3)
```sh
clang -O3 main.c -o main3
objdump -d main3
```
최적화로 인해 `calculation`을 호출하지 않고 바로 90으로 inline 됨. (`mov $0x5a,%esi`)
```text
0000000000001140 <doubleNum>:
    1140:       8d 04 3f                lea    (%rdi,%rdi,1),%eax
    1143:       c3                      ret    
    1144:       66 2e 0f 1f 84 00 00    cs nopw 0x0(%rax,%rax,1)
    114b:       00 00 00 
    114e:       66 90                   xchg   %ax,%ax

0000000000001150 <square>:
    1150:       89 f8                   mov    %edi,%eax
    1152:       0f af c7                imul   %edi,%eax
    1155:       c3                      ret    
    1156:       66 2e 0f 1f 84 00 00    cs nopw 0x0(%rax,%rax,1)
    115d:       00 00 00 

0000000000001160 <calculation>:
    1160:       0f af ff                imul   %edi,%edi
    1163:       0f af fe                imul   %esi,%edi
    1166:       8d 04 3f                lea    (%rdi,%rdi,1),%eax
    1169:       c3                      ret    
    116a:       66 0f 1f 44 00 00       nopw   0x0(%rax,%rax,1)

0000000000001170 <main>:
    1170:       50                      push   %rax
    1171:       48 8d 3d 8c 0e 00 00    lea    0xe8c(%rip),%rdi        # 2004 <_IO_stdin_used+0x4>
    1178:       be 5a 00 00 00          mov    $0x5a,%esi
    117d:       31 c0                   xor    %eax,%eax
    117f:       e8 ac fe ff ff          call   1030 <printf@plt>
    1184:       b8 5a 00 00 00          mov    $0x5a,%eax
    1189:       59                      pop    %rcx
    118a:       c3                      ret    

```

## All optimization wihtout violating strict compliance (-O3) - Separating code into several files
아래처럼 파일을 분리해놓은 경우 inline 최적화가 안되었는데 컴파일러가 참조없이 각각의 source file에 대해 object file을 만들기 때문.

```sh
clang -O3 main1.c calc.c ops.c -o main3_2
objdump -d main3_2
```
```text
0000000000001140 <main>:
    1140:       53                      push   %rbx
    1141:       bf 03 00 00 00          mov    $0x3,%edi
    1146:       be 05 00 00 00          mov    $0x5,%esi
    114b:       e8 20 00 00 00          call   1170 <calculation>
    1150:       89 c3                   mov    %eax,%ebx
    1152:       48 8d 3d ab 0e 00 00    lea    0xeab(%rip),%rdi        # 2004 <_IO_stdin_used+0x4>
    1159:       89 c6                   mov    %eax,%esi
    115b:       31 c0                   xor    %eax,%eax
    115d:       e8 ce fe ff ff          call   1030 <printf@plt>
    1162:       89 d8                   mov    %ebx,%eax
    1164:       5b                      pop    %rbx
    1165:       c3                      ret    
    1166:       66 2e 0f 1f 84 00 00    cs nopw 0x0(%rax,%rax,1)
    116d:       00 00 00 

0000000000001170 <calculation>:
    1170:       55                      push   %rbp
    1171:       53                      push   %rbx
    1172:       50                      push   %rax
    1173:       89 f3                   mov    %esi,%ebx
    1175:       e8 26 00 00 00          call   11a0 <square>
    117a:       89 c5                   mov    %eax,%ebp
    117c:       89 df                   mov    %ebx,%edi
    117e:       e8 0d 00 00 00          call   1190 <doubleNum>
    1183:       0f af c5                imul   %ebp,%eax
    1186:       48 83 c4 08             add    $0x8,%rsp
    118a:       5b                      pop    %rbx
    118b:       5d                      pop    %rbp
    118c:       c3                      ret    
    118d:       0f 1f 00                nopl   (%rax)

0000000000001190 <doubleNum>:
    1190:       8d 04 3f                lea    (%rdi,%rdi,1),%eax
    1193:       c3                      ret    
    1194:       66 2e 0f 1f 84 00 00    cs nopw 0x0(%rax,%rax,1)
    119b:       00 00 00 
    119e:       66 90                   xchg   %ax,%ax

00000000000011a0 <square>:
    11a0:       89 f8                   mov    %edi,%eax
    11a2:       0f af c7                imul   %edi,%eax
    11a5:       c3                      ret    
```

# Link-Time Optimization (LTO)
Compiler Optimization의 3번째 케이스같은 상황에서 최적화를 위해 LTO 필요. 아래 보면 다시 `mov $0x5a,%esi`로 inline된 걸 볼 수 있음.
```sh
clang -O3 -flto main1.c calc.c ops.c -o main3_3
```
```text
0000000000001140 <main>:
    1140:       50                      push   %rax
    1141:       48 8d 3d bc 0e 00 00    lea    0xebc(%rip),%rdi        # 2004 <_IO_stdin_used+0x4>
    1148:       be 5a 00 00 00          mov    $0x5a,%esi
    114d:       31 c0                   xor    %eax,%eax
    114f:       e8 dc fe ff ff          call   1030 <printf@plt>
    1154:       b8 5a 00 00 00          mov    $0x5a,%eax
    1159:       59                      pop    %rcx
    115a:       c3                      ret    
```