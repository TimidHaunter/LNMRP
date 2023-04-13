<?php

##闭包的使用
$a = function () {
    echo '我是匿名函数' . PHP_EOL;
};
$a();

/**
 * 闭包函数当做参数传入到函数中
 * @param $a
 * @return void
 */
function testA($a)
{
    var_dump($a);
}

testA($a);


$b = function ($name) {
    echo '匿名函数的名字是:' . $name . PHP_EOL;
};
$b('无敌匿名函数');

##use使用
$age = 16;
$c = function ($name) {
    echo '闭包传入的参数:' . $name . ', 闭包没有传入的参数:' . $age . PHP_EOL;
};

$c('我是闭包');
/**
 * 函数体内部不能调用外部$age变量，要使用use传递到函数体内
 */

$age = 16;
$c = function ($name) use ($age) {
    echo '闭包传入的参数:' . $name . ', 闭包没有传入的参数:' . $age . PHP_EOL;
};

$c('我是使用了use的闭包');

echo '=================================================================' . PHP_EOL;

##作用域
/**
 * 普通函数
 * @return void
 */
function globalVar()
{
    global $outVar;
    echo $outVar . PHP_EOL;
}

/**
 * 闭包函数
 * @return void
 */
$d = function () use ($outVar) {
    echo $outVar . PHP_EOL;
};

/**
 * 匿名函数
 */
$e = function () {
    global $outVar;
    echo $outVar . PHP_EOL;
};

$outVar = '我是外部变量1';
$d();
globalVar();
$e();

$outVar = '我是外部变量2';
$f = function () use ($outVar) {
    echo $outVar . PHP_EOL;
};
$f();

##值传递还是引用传递
$outVar = '我是外部变量,我修改了值';
$f(); //我是外部变量2,值传递

$outVar = '我是外部变量,我修改了值,使用了&引用传递';
$f = function () use (&$outVar) {
    echo $outVar . PHP_EOL;
};
$f(); //我是外部变量,我修改了值,使用了&引用传递

$outVar = '我是外部变量,我修改了值2,使用了&引用传递';
$f(); //我是外部变量,我修改了值2,使用了&引用传递

echo '=================================================================' . PHP_EOL;

/**
 * 闭包的应用
 */
$arr = [
    ['name' => '李白'],
    ['name' => '杜甫'],
    ['name' => '苏轼']
];

$string = '是一个大文豪!';

foreach ($arr as $k => $v) {
    $arr[$k]['name'] = $v['name'] . $string;
}
var_dump($arr);

/**
 * 针对数组每一项进行回调操作
 * 类似的还有array_map
 */
array_walk($arr, function (&$v) use ($string) {
    $v['name'] .= '还是一个古人!';
});
var_dump($arr);

/**
 * 普通函数
 */
function testCommon()
{
    //匿名函数
    return function ($name) {
        echo '我是普通函数中的匿名函数入参:' . $name . PHP_EOL;
    };
}

//调用普通函数中的匿名函数
testCommon()('好牛逼的用法');

echo '=================================================================' . PHP_EOL;

class A
{}

class B
{}

class C
{}

class Ioc
{
    //类
    public $objs = [];
    public $containers = [];

    public function __construct()
    {
        //闭包作用，延迟加载，节省内存
        //如果没有闭包
//        $this->objs['A'] = new A();
//        $this->objs['B'] = new A();
//        $this->objs['C'] = new A();

        $this->objs['A'] = function () {
            return new A();
        };

        $this->objs['B'] = function () {
            return new B();
        };

        $this->objs['C'] = function () {
            return new C();
        };
    }

    public function bind($name)
    {
        if (!isset($this->containers[$name])) {
//            var_dump($this->objs[$name], isset($this->objs[$name]));
            if (isset($this->objs[$name])) {
                $this->containers[$name] = $this->objs[$name]();
            } else {
                return null;
            }
        }
        return $this->containers[$name];
    }
}

$ioc = new Ioc();
var_dump($ioc);
$bClass = $ioc->bind('A');
$cClass = $ioc->bind('B');
$dClass = $ioc->bind('C');
$eClass = $ioc->bind('D');

var_dump($bClass); // B
var_dump($cClass); // C
var_dump($dClass); // D
var_dump($eClass); // NULL