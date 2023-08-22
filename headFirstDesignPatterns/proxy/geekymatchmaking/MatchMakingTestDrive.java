public class MatchMakingTestDrive {
  private MyProxy myProxy;

  public static void main(String[] args) {
    MatchMakingTestDrive test = new MatchMakingTestDrive();
    test.drive();
  }

  public MatchMakingTestDrive() {
    this.myProxy = new MyProxy();
    initializeDatabase();
  }

  private void initializeDatabase() {

  }

  public void drive() {
    Person joe = getPersonFromDatabase("Joe Javabean");
    Person ownerProxy = getOwnerProxy(joe);
    System.out.printf("Name is %s%n", ownerProxy.getName());
    ownerProxy.setInterest("bowling, Go");
    System.out.println("Interests set from owner proxy");
    try {
      ownerProxy.setGeekRating(10);
    } catch (Exception e) {
      System.out.println("Can't set rating from owner proxy");
    }
    System.out.printf("Rating is %d%n", ownerProxy.getGeekRating());

    Person nonOwnerProxy = getNonOwnerProxy(joe);
    System.out.printf("Name is %s%n", nonOwnerProxy.getName());
    try {
      nonOwnerProxy.setInterest("bowling, Go");
    } catch (Exception e) {
      System.out.println("Can't set interests from non owner proxy");
    }
    nonOwnerProxy.setGeekRating(10);
    System.out.println("Rating set from non owner proxy");
    System.out.printf("Rating is %d%n", nonOwnerProxy.getGeekRating());

  }

  private Person getPersonFromDatabase(String name) {
    Person person = new PersonImpl();
    person.setName(name);
    person.setGeekRating(5);
    person.setGender("m");
    return person;
  }

  private Person getOwnerProxy(Person person) {
    return myProxy.getOwnerProxy(person);
  }

  private Person getNonOwnerProxy(Person person) {
    return myProxy.getNonOwnerProxy(person);
  }

}
