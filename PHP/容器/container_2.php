<?php

/**
 * https://blog.csdn.net/EasyTure/article/details/112603321
 *
 * ReflectionClass
 * ReflectionObject
 * 反射时的区别
 */

//class test {
//    private $name;
//    private $sex;
//
//    function __construct(){
//        $this->aaa='aaa';
//    }
//}
//
//$test = new test();
//
//// 反射类声明时的结构
//$reflect = new ReflectionClass($test);
//
//// 反射实例化后的结构
//$reflect = new ReflectionObject($test);
//
//// 通过对象反射得到类的属性
//$pro = $reflect->getDefaultProperties();
//var_dump($pro);
//
//// 打印结果：aaa
//echo $test->aaa;


class Point
{
    public $x;
    public $y;

    /**
     * Point constructor.
     * @param int $x  horizontal value of point's coordinate
     * @param int $y  vertical value of point's coordinate
     */
    public function __construct($x = 0, $y = 0)
    {
        $this->x = $x;
        $this->y = $y;
    }
}

class Circle
{
    /**
     * @var int
     * 半径
     */
    public $radius;

    /**
     * @var Point
     * 圆心点
     */
    public $center;

    const PI = 3.14;

    /**
     * Circle constructor.
     * @param Point $point
     * Circle依赖于Point类
     * @param int $radius
     */
    public function __construct(Point $point, $radius = 1)
    {
        $this->center = $point;
        $this->radius = $radius;
    }

    // 打印圆点的坐标
    public function printCenter()
    {
        printf('center coordinate is (%d, %d)', $this->center->x, $this->center->y);
    }

    // 计算圆形的面积
    public function area()
    {
        return 3.14 * pow($this->radius, 2);
    }
}

// 构建类的对象
function make($className)
{
    $reflectionClass = new ReflectionClass($className);
    // 获得构造函数结构
    $constructor = $reflectionClass->getConstructor();
    // 获得构造函数参数
    $parameters  = $constructor->getParameters();
    // 依赖解析，因为依赖都是通过参数的形式传递进来的
    $dependencies = getDependencies($parameters);

    return $reflectionClass->newInstanceArgs($dependencies);
}

$circle = make('Circle');
$area = $circle->area();

// 依赖解析
function getDependencies($parameters)
{
    $dependencies = [];
//    var_dump($parameters);
    foreach($parameters as $parameter) {
//        var_dump($parameter);
        // 通过依赖关系，找到对应的类
        $dependency = $parameter->getClass();
//        var_dump($dependency);

        if (is_null($dependency)) {
            // 如果找不到对应的类，说明依赖关系没有了
            if ($parameter->isDefaultValueAvailable()) {
                $dependencies[] = $parameter->getDefaultValue();
            } else {
                //不是可选参数的为了简单直接赋值为字符串0
                //针对构造方法的必须参数这个情况
                //laravel是通过service provider注册closure到IocContainer,
                //在closure里可以通过return new Class($param1, $param2)来返回类的实例
                //然后在make时回调这个closure即可解析出对象
                //具体细节我会在另一篇文章里面描述
                $dependencies[] = '0';
            }
        } else {
            // 递归解析出依赖类的对象，point->x,y radius
            $dependencies[] = make($parameter->getClass()->name);
        }
    }

//    var_dump($dependencies);
    return $dependencies;
}

//$point = new Point(2, 4);
//$circle = new Circle($point);
//
//$reflectionClass = new reflectionClass($circle);
//$reflectionClass = new reflectionClass('Circle');

$reflectionClass = new reflectionClass(Circle::class);

//var_dump($reflectionClass);
//object(ReflectionClass)#3 (1) {
//  ["name"]=>
//  string(6) "Circle"
//}

// 常量
$constants = $reflectionClass->getConstants();
//var_dump($constants);
//array(1) {
//  ["PI"]=>
//  float(3.14)
//}

// 类的方法
$properties = $reflectionClass->getProperties();
//var_dump($properties);
//array(2) {
//  [0]=>
//  object(ReflectionProperty)#4 (2) {
//    ["name"]=>
//    string(6) "radius"
//    ["class"]=>
//    string(6) "Circle"
//  }
//  [1]=>
//  object(ReflectionProperty)#5 (2) {
//    ["name"]=>
//    string(6) "center"
//    ["class"]=>
//    string(6) "Circle"
//  }
//}

// 反射出类中定义的方法
$methods = $reflectionClass->getMethods();

//$refFunction = new ReflectionFunction('area');
//$parameters = $refFunction->getParameters();
//var_dump($parameters);

// 单独获取类的构造方法
$constructor = $reflectionClass->getConstructor();

// 反射出方法的参数，这里特指构造函数方法参数
$parameters = $constructor->getParameters();
//var_dump($parameters);