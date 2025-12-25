package com.example;

import reactor.core.publisher.Flux;
import reactor.core.publisher.Sinks;

/**
 * @author wangxiang
 * @description
 * @create 2025/12/25 21:55
 */
public class ReactorBasics {
    public static void main(String[] args) throws InterruptedException {
        // === 示例1: 理解Flux ===
        System.out.println("=== Flux示例 ===");
        Flux<String> flux = Flux.just("苹果", "香蕉", "橙子");

        flux.subscribe(
                fruit -> System.out.println("收到水果: " + fruit),  // 处理每个元素
                error -> System.err.println("出错了: " + error),    // 处理错误
                () -> System.out.println("水果发完了!")              // 完成信号
        );

        // === 示例2: 理解Sinks.Many ===
        // Sinks.Many就像一个"可控制的水龙头"
        Sinks.Many<String> sink = Sinks.many()
                .multicast()
                .onBackpressureBuffer();

        // 转成Flux给别人订阅
        Flux<String> stream = sink.asFlux();

        // 订阅者1
        stream.subscribe(data ->
                System.out.println("订阅者1收到: " + data)
        );

        // 手动"拧开水龙头"发送数据
        sink.tryEmitNext("第一条消息");
        sink.tryEmitNext("第二条消息");

        // 订阅者2(晚加入,只能收到后续消息)
        stream.subscribe(data ->
                System.out.println("订阅者2收到: " + data)
        );

        sink.tryEmitNext("第三条消息");
        sink.tryEmitComplete(); // 关闭水龙头

        Thread.sleep(1000); // 等待异步处理完成
    }
}
