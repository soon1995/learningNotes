package com.alsoon.studentms.bean;

import java.util.ArrayList;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class UserEditVo {
  private Integer u_id;
  private String u_name;
  private String u_pwd;
  private String u_phone;
  private Integer u_role;
  private Integer u_isdeleted;
  // private ArrayList<String> c_ids;
  private String[] c_ids;
  private String[] c_names;
  private String[] c_scores;
  
}
