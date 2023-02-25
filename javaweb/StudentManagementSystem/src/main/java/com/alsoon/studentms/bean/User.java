package com.alsoon.studentms.bean;

import java.io.Serializable;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class User implements Serializable {
  private Integer u_id;
  private String u_name;
  private String u_pwd;
  private String u_phone;
  private Integer u_role;
  private Integer u_isdeleted;
}
