package com.wwxiong.singleton.type1;

// ���δ���ʵ���˶���ʽ����ģʽ��
// ����ģʽ��һ�ִ��������ģʽ����ȷ��һ����ֻ��һ��ʵ�������ṩ�Ը�ʵ����ȫ�ַ��ʵ㡣

public class SingletonTest01 {

	public static void main(String[] args) {
		// ���� Singleton ���ʵ����ʹ�� getInstance() ����
		Singleton instance = Singleton.getInstance();
		Singleton instance2 = Singleton.getInstance();
		// �Ƚ�����ʵ�������Ƿ����
		System.out.println(instance == instance2); // true
		System.out.println("instance.hashCode=" + instance.hashCode());
		System.out.println("instance2.hashCode=" + instance2.hashCode());
	}

}

// Singleton ���ǵ���ģʽ��ʵ���࣬���������ص㣺
// ������˽�л����ⲿ�޷�ͨ�� new �ؼ��ִ���ʵ����ֻ��ͨ�� getInstance() ������ȡ������
// �����ڲ�����һ��˽�о�̬���� instance������ Singleton ���Ψһʵ�����������ʱ�ʹ����˸�ʵ������˱���Ϊ������ʽ����
// �ṩһ�����еľ�̬���� getInstance()���÷������� Singleton ���Ψһʵ�� instance��

class Singleton {
	
	// ������˽�л����ⲿ�޷�ͨ�� new �ؼ��ִ���ʵ��
	private Singleton() {
		
	}
	
	// �����ڲ�����һ��˽�о�̬���� instance������ Singleton ���Ψһʵ��
	private final static Singleton instance = new Singleton();
	
	// �ṩһ�����еľ�̬���� getInstance()������ʵ������
	public static Singleton getInstance() {
		return instance;
	}
	
}