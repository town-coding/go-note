package com.example;

import reactor.core.publisher.Flux;

/**
 * @author wangxiang
 * @description
 * @create 2025/12/25 22:13
 */
public class FluxTest {
    public static void main(String[] args) {
        Flux<Integer> generated = Flux.generate(
                () -> 0,  // 初始状态
                (state, sink) -> {
                    sink.next(state);  // 发射当前状态
                    if (state == 10) {
                        sink.complete();  // 达到条件则完成
                    }
                    return state + 1;  // 更新状态
                }
        );
        generated.subscribe(System.out::println);
    }
}
