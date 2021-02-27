# Recursion

## 재귀 호출의 일반적인 형태

```python
# 일반적인 재귀의 형태 1
def function(입력):
    if 입력 > 일정값: # 입력이 일정 값 이상이면
        return function(입력-1) # 입력보다 작은 값
    return 일정값 # 재귀 호출 종료
```
```python
# 일반적인 재귀의 형태 2
def function(입력):
    if 입력 <= 일정값: # 입력이 일정 값보다 작으면
        return 결과값 # 재귀 호출 종료
    function(입력보다 작은 값)
    return 결과값
```

```go
// 일반적인 재귀의 형태 1
func function(입력 타입) 리턴 타입 {
    if 입력 > 일정값 {
        return function(입력-1)
    }
    return 일정값
}
```

```go
// 일반적인 재귀의 형태 2
func function(입력 타입) 리턴 타입 {
    if 입력 <= 일정값 {
        return 결과값
    }
    function(입력보다 작은 값)
    return 결과값
}
```