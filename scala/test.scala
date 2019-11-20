// https://www.cnblogs.com/lq0310/p/9840317.html
val s="hello world"
println(s)

val data = List(3,4,2,1,5,"hello")
for (s<-data) println(s)

val info = Array(2,3,4,5,6,7,7,8,89,1)
for (ss<-info) println(ss)

println("######")
info.foreach(println)

// java系有高人
class Student {
    private var Name:String = "Tome"
    private var stuAge:Int = 20

    def getName():String = Name
    def setName(name:String) = this.Name = name

    def getAge():Int = stuAge
    def setAge(newAge:Int) = this.stuAge=newAge
}

var stu = new Student

println(stu.getName())

// trait 特质 多重继承
trait Human {
    val id:Int
    val name:String

    def sayHello():String = "hello "+name
}

trait Actions {
    def getActionNames():String
}

class Teacher (val id:Int,val name:String) extends Human with Actions {
    def getActionNames():String = "Action is running"
}

object Demo2 {
    def main(args: Array[String]) {
        val s1 = new Teacher(1,"Lixueping")
        println(s1.sayHello())
        println(s1.getActionNames())
    }
}

// Demo2()

val s1 = new Teacher(1,"Lixueping")
println(s1.sayHello())
println(s1.getActionNames())