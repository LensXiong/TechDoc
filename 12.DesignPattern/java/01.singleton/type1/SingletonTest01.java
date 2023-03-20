package com.wwxiong.singleton.type1;

// 本段代码实现了饿汉式单例模式。
// 单例模式是一种创建型设计模式，它确保一个类只有一个实例，并提供对该实例的全局访问点。

public class SingletonTest01 {

	public static void main(String[] args) {
		// 创建 Singleton 类的实例并使用 getInstance() 方法
		Singleton instance = Singleton.getInstance();
		Singleton instance2 = Singleton.getInstance();
		// 比较两个实例引用是否相等
		System.out.println(instance == instance2); // true
		System.out.println("instance.hashCode=" + instance.hashCode());
		System.out.println("instance2.hashCode=" + instance2.hashCode());
	}

}

// Singleton 类是单例模式的实现类，它有以下特点：
// 构造器私有化，外部无法通过 new 关键字创建实例，只能通过 getInstance() 方法获取单例。
// 本类内部创建一个私有静态变量 instance，它是 Singleton 类的唯一实例。在类加载时就创建了该实例，因此被称为“饿汉式”。
// 提供一个公有的静态方法 getInstance()，该方法返回 Singleton 类的唯一实例 instance。

class Singleton {
	
	// 构造器私有化，外部无法通过 new 关键字创建实例
	private Singleton() {
		
	}
	
	// 本类内部创建一个私有静态变量 instance，它是 Singleton 类的唯一实例
	private final static Singleton instance = new Singleton();
	
	// 提供一个公有的静态方法 getInstance()，返回实例对象
	public static Singleton getInstance() {
		return instance;
	}
	
}