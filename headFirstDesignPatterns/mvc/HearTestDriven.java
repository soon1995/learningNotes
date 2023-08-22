public class HearTestDriven {
  public static void main(String[] args) {
    HeartModel heartModel = new HeartModel();
    ControllerInterface model = new HeartController(heartModel);
  }
}
