import akka.actor._

class Actor1 extends Actor {
    override def act():Unit = {
        for(i<- 1 to 10) {
            println("Actor1 ======"+i)
        }
    }
}

object Actor2 extends Actor {
    override def act():Unit = {
        for(j<- 1 to 10) {
            println("Actor2===="+j)
        }
    }
}

object Actor1 {
    def main(args: Array[String]):Unit = {
        val actor = new Actor1
        actor.act()
        Actor2.act()
    }
}