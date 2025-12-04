#include <stdio.h>
#include <iostream>
#include <math.h>

using namespace std;

struct node {
    long long value;
    node *next;
} *qHead, *rHead;

node *createNode(long long value) {
    node *newNode = (node *)malloc(sizeof(node));
    newNode->value = value;
    newNode->next = NULL;
    return newNode;
}

void addOrdered(node **head, long long value) {
    node *newNode = createNode(value);
    if (*head == NULL) {
        *head = newNode;
    } else {
        node *current = *head;
        while (current->next != NULL && current->next->value < value) {
            current = current->next;
        }
    }
    newNode->next = current->next;
    current->next = newNode;
}

void sortArray(long long *a, int start, int end, bool useRand) {
    if (start >= end) return;

    int middle = (useRand) ? floor(rand() % (end - start + 1) + start) : floor((start + end) / 2);

    sortArray(a, start, middle, useRand);
    sortArray(a, middle + 1, end, useRand);

    int lenLeft = middle - start + 1;
    int lenRight = end - middle;

    long long tmpLeft[lenLeft] = {0};
    long long tmpRight[lenRight] = {0};

    for (int i = start; i <= middle; i++) tmpLeft[i - start] = a[i];
    for (int i = middle + 1; i <= end; i++) tmpRight[i - middle - 1] = a[i];

    int iL = 0, iR = 0;

    for (int i = 0; i <= end - start; i++) {
        if (iL >= lenLeft) {
            a[start + i] = tmpRight[iR++];
            continue;
        }    
        
        if (iR >= lenRight) {
            a[start + i] = tmpLeft[iL++];
            continue;
        }

        if (tmpLeft[iL] < tmpRight[iR])
            a[start + i] = tmpLeft[iL++];
        else
            a[start + i] = tmpRight[iR++];
    }
}

int main(int argc, char **argv) {
    int t;
    scanf("%d", &t);
    for (int i = 0; i < t; i++) {
        int n;
        long long k;
        scanf("%d", &n);
        scanf("%lld", &k);

        long long q[n], r[n], qT[n];
        for (int j = 0; j < n; j++) {
            scanf("%lld", q + j);
        }
        for (int j = 0; j < n; j++) {
            qT[j] = (k - q[j]) / (1 + q[j]);
        }
        for (int j = 0; j < n; j++) scanf("%lld", r + j);

        sortArray(qT, 0, n-1, true);
        sortArray(r, 0, n-1, true);

        int matchCount = 0, itQ = 0, itR = 0, itLastR = 0;
        while (itQ < n && itR < n) {
            if (r[itR] <= qT[itQ]) { matchCount++; itR++; itLastR = itR; itQ++; }
            else itR++;

            if (itR == n) { itR = itLastR; itQ++; }
        }

        printf("%d\n", matchCount);
    }

}