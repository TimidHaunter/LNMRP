<?php

# https://www.bilibili.com/video/BV1ko4y1R7mz?spm_id_from=333.999.0.0
# https://blog.csdn.net/qq_20124631/article/details/113456378?spm=1001.2014.3001.5502
#
# 1.老式类之间，需要在类中new别的类，强依赖
# 2.把new对象的过程放在类外，把new好的对象，以参数形式传递到类中，实现依赖注入
# 3.引入容器的概念，所有的new对象都有容器完成，不需要再手动new，手动传参

/**
 * 老式类之间的关系
 * phone 类依赖于 usb 类，因为在 phone 类构造函数里需要主动 new usb
 */
class phone
{
    public $obj;

    public function __construct($obj)
    {
        // 1.老式依赖
        // 注意这里主动去 new 了需要的对象
//        $this->obj = new usb();

        // 2.改进一下
        // 通过变量，把new usb传参进来
        $this->obj = $obj;
    }

    public function act()
    {
        $this->obj->doSomething();
    }
}

class usb
{
    public function doSomething(){
        echo 'usb_doSomething'.PHP_EOL;
    }
}


class battery
{
    public function doSomething(){
        echo 'battery_doSomething'.PHP_EOL;
    }
}

// 2.改进
// 依赖注入，控制反转

// 3.有了 container 就不需要在手动 new 对象了，全部交给容器
//$usb = new usb;
//
//$phone = new phone($usb);
//$phone->act();

// 上面都是靠我们自己手动反转，自己需要new 对象




/**
 * 这个时候就需要容器去自动 new 对象
 * 3.引入容器类
 */
class container
{
    // 存储资源
    public $objs = [];

    // 绑定关系，把主要的类绑定关系
    public function bind($class, Closure $closure_obj)
    {
        // 判断 $closure_obj 参数是不是一个闭包
        if (! $closure_obj instanceof Closure){
            // 不是闭包，想办法变成一个闭包
            // 靠反射变成闭包
            $this->objs[$class] = function ()use($closure_obj){
                return $this->ref($closure_obj);
            };
        }
        // 是闭包，直接赋值给$this->objs[$class]
        else {
            $this->objs[$class] = $closure_obj;
        }
    }

    // 在需要的该类的地方取出关系
    public function make($class)
    {
        // 容器中是否有绑定的关系
        if (isset($this->objs[$class])) {
            $new = $this->objs[$class];
        }
        // 不在绑定关系中，需要反射
        else {
            $new = $this->ref($class);
        }

        return $new();
    }

    /**
     * 4.引入反射机制
     * 反射的机制明白了，就是不明白为什么用反射来优化容器
     */
    public function ref($class, $params=[]) {
        // 通过对象反射出类声明时的结构
        $ref = new ReflectionClass($class);

        // 判断这个类能不能被实例化
        if ($ref->isInstantiable()) {
            // 获得类的构造函数
            $constructor = $ref->getConstructor();

            // 判断类是否有构造函数
            if (!is_null($constructor)){
                // 获得构造函数的参数，也就是依赖关系
                $pars = $constructor->getParameters();

                // 如果参数是空，就是没有依赖关系，直接返回对象
                if (empty($pars)) {
                    return new $class;
                }
                // 如果构造函数参数不为空，有依赖关系
                else {
                    // 循环依赖关系
                    foreach ($pars as $par) {
                        $dependencyClass = $par->getClass();
                        if (is_null($dependencyClass)) {
                            $dependencies[] = NULL;
                        } else {
                            // 类存在创建类实例
                            $dependencies[] = $this->make($par->getClass()->name);
                        }
                    }
                    // 返回一个对象 new class_name()
                    return $ref->newInstanceArgs($dependencies);
                }
            }
            // 没有构造函数，直接返回对象
            else
            {
                return new $class;
            }
        }
        // 如果不能被实例化，就不能拿来当对象用，直接返回null
        return null;
    }
}


/**
 * 3.容器的概念
 */
$container = new container;
// 用绑定的方法，将usb类对象绑定到容器
$container->bind('usb', function(){
    return new usb;
});
// phone类需要usb类对象的时候取出关系
$phone = new phone($container->make('usb'));
/**
 * 3.end
 */