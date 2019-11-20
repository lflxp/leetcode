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