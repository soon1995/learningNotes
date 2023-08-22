import java.io.File;
import java.util.ArrayList;
import java.util.List;

import javax.sound.sampled.AudioSystem;
import javax.sound.sampled.Clip;
import javax.sound.sampled.Line;

public class BeatModel implements BeatModelInterface, Runnable {
  List<BeatObserver> beatObservers = new ArrayList<BeatObserver>();
  List<BPMObserver> bpmObservers = new ArrayList<BPMObserver>();
  int bpm = 90;
  Thread thread;
  boolean stop = false;
  Clip clip;

  public void initialize() {
    try {
      File resource = new File("sample.wav");
      clip = (Clip) AudioSystem.getLine(new Line.Info(Clip.class));
      clip.open(AudioSystem.getAudioInputStream(resource));
    } catch (Exception e) {
      // TODO: handle exception
      e.printStackTrace();
    }
  }

  public void on() {
    bpm = 90;
    notifyBPMObservers();
    thread = new Thread(this);
    stop = false;
    thread.start();
  }

  public void off() {
    stopBeat();
    stop = true;
  }

  public void run() {
    while (!stop) {
      playBeat();
      notifyBeatObservers();
      try {
        Thread.sleep(60000 / getBPM());
        System.out.println("a");
      } catch (Exception e) {
        e.printStackTrace();
      }
    }
  }

  public void setBPM(int bpm) {
    this.bpm = bpm;
    notifyBPMObservers();
  }

  public int getBPM() {
    return bpm;
  }

  public void registerObserver(BeatObserver o) {
    beatObservers.add(o);
  }

  public void registerObserver(BPMObserver o) {
    bpmObservers.add(o);
  }

  public void playBeat() {
    clip.setFramePosition(0);
    clip.start();
  }

  public void stopBeat() {
    clip.setFramePosition(0);
    clip.stop();
  }

  private void notifyBPMObservers() {
    bpmObservers.forEach(o -> {
      o.updateBPM();
    });
  }

  private void notifyBeatObservers() {
    beatObservers.forEach(o -> {
      o.updateBeat();
    });
  }

  public void removeObserver(BeatObserver o) {
    int i = beatObservers.indexOf(o);
    if (i >= 0) {
      beatObservers.remove(i);
    }
  }

  public void removeObserver(BPMObserver o) {
    int i = bpmObservers.indexOf(o);
    if (i >= 0) {
      bpmObservers.remove(i);
    }
  }

}
