package _1_singleton

import "fmt"

// 饿汉式（静态变量）
type Singleton struct{}

// 在包加载时，初始化一个唯一的 Singleton 实例
// 使用了一个全局变量 instance 来存储 Singleton 类的唯一实例。
// 在包加载时，Go 会自动调用 init 函数，可以在 init 函数中初始化这个实例。
var instance *Singleton = &Singleton{}

// 提供一个公有的静态方法，返回实例对象
func GetInstance() *Singleton {
	return instance
}

// 饿汉式单例模式的主要问题有两个：
// ① 在类加载时就创建了实例，如果该实例在后续程序运行中没有被使用到，就会造成内存浪费。
// ② 在多线程环境中，饿汉式单例模式可能会存在线程安全问题。

// 在多线程环境中，如果多个线程同时调用饿汉式单例模式的 getInstance() 方法，可能会存在线程安全问题，主要原因是：
// ① 竞态条件：当多个线程同时调用 getInstance() 方法时，可能会出现竞态条件，即多个线程同时读取 instance 变量的值，
// 如果此时 instance 变量还没有被初始化，那么就可能会创建多个实例。
// ② 可见性问题：当一个线程创建了 Singleton 实例后，其他线程可能无法立即看到该变化，因为内存缓存和指令重排等机制可能会导致线程之间的可见性问题。
// 即一个线程对 instance 变量的修改可能无法立即被其他线程感知，从而导致其他线程创建新的实例。
// 为了解决这些问题，我们需要使用线程安全的单例模式实现方式，如使用懒汉式单例模式或者双重检查锁定单例模式等。
func main() {
	// 测试
	instance := GetInstance()
	instance2 := GetInstance()

	fmt.Println(instance == instance2) // true
	fmt.Println("instance:", instance)
	fmt.Println("instance2:", instance2)
}
