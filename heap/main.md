# Heap
## 힙(Heap) 이란?
* 힙 : 데이터에서 최댓값과 최솟값을 빠르게 찾기 위해 고안된 완전 이진 트리(Complete Binary Tree)
    * 완전 이진 트리 : 노드를 삽입할 때 최하단 왼쪽 노드부터 차례대로 삽입하는 트리

* 힙을 사용하는 이유
    - 배열에 데이터를 넣고, 최댓값과 최솟값을 찾으려면 `O(n)` 이 걸림
    - 이에 반해, 힙에 데이터를 넣고, 최댓값과 최솟값을 찾으려면 `O(logn)` 이 걸림
    - **우선순위 큐**와 같이 최댓값 또는 최솟값을 빠르게 찾아야 하는 자료구조 및 알고리즘의 구현 등에 활용됨

## 힙(Heap) 구조
* 힙은 최댓값을 구하기 우한 구조(최대 힙, Max Heap)와, 최솟값을 구하기 위한 구조(최소 힙, Min Heap)로 분류할 수 있음
* 힙은 다음과 같이 두 가지 조건을 가지고 있는 자료구조임
    1. 각 노드의 값은 해당 노드의 자식 노드가 가진 값보다 크거나 같다. (최대 힙의 경우)
        * 최소 힙의 경우는 각 노드의 값은 해당 노드의 자식 노드가 가진 값보다 같거나 작음.
    2. 완전 이진 트리 형태를 가짐

## 힙과 이진 탐색 트리의 공통점과 차이점
* 공통점 : 힙과 이진 탐색 트리는 모두 이진 트리임
* 차이점 :
    * 힙은 각 노드의 값이 자식 노드보다 크거나 같음(Max Heap)
    * 이진 탐색 트리는 왼쪽 자식 노드의 값이 가장 작고, 그 다음 부모 노드, 그 다음 오른쪽 자식 노드 값이 가장 틈
    * 힙은 이진 탐색 트리의 조건인 자식 노드에서 작은 값은 왼쪽, 큰 값은 오른쪽이라는 조건은 없음
        * 힙의 왼쪽 및 오른쪽 자식 노드의 값은 오른쪽이 클 수도 있고, 왼쪽이 클 수도 있음
    * 이진 탐색 트리는 탐색을 위한 구조, 힙은 최댓/최솟값 검색을 위한 구조 중 하나로 이해하면 됨

## 힙에 데이터 추가하기(Max Heap)
1. 최하단 왼쪽부터 차례대로 데이터를 확인한다. 데이터가 nil인 상태가 발견되면, 해당 위치에 데이터를 추가한다. 
2. 해당 위치의 부모 노드와 자신의 값을 비교해서 자신이 더 크면, 부모 노드와 swap한다. 
3. 2번 과정을 부모 노드가 더 커질 때까지, 또는 root node일 때까지 반복한다.

## 힙에서 데이터 삭제하기(Max Heap)
- 보통 삭제는 최상단 노드(root 노드)를 삭제하는 것이 일반적임
    * 힙의 용도는 최댓값 또는 최솟값을 root 노드에 놓아서, 최댓값과 최솟값을 바로 꺼내 쓸 수 있도록 하는 것임.
- 상단의 데이터 삭제시, 가장 최하단부 왼쪽에 위치한 노드(**일반적으로 가장 마지막에 추가한 노드**)를 root 노드로 이동
- root 노드의 값이 child node 보다 작을 경우, root 노드의 child node 중 가장 큰 값을 가진 노드와 root 노드의 위치를 바꿔주는 작업을 반복함(swap)

## 힙 구현
### 힙과 배열
- 일반적으로 힙 구현 시, 배열 자료구조를 활용함
- 배열은 인덱스가 0번부터 시작하지만, 힙 구현의 편의를 위해, root 노드 인덱스 번호를 1로 지정하면 구현이 좀더 수월함.
    - 부모 노드 인덱스 번호 = 자식 노드 인덱스 번호 // 2
    - 왼쪽 자식 노드 인덱스 번호 = 부모 노드 인덱스 번호 * 2 
    - 오른쪽 자식 노드 인덱스 번호 = 부모 노드 인덱스 번호 * 2 + 1