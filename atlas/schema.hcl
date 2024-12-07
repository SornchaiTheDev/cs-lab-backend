schema "public" {
  comment = "standard public schema"
}

enum "role" {
  schema = schema.public
  values = [ "student" , "instructor" , "admin" ]
}

table "users" {
  schema = schema.public
  column "id" {
    type = uuid
    default = sql("gen_random_uuid()")
  }
  column "email" {
    type = text
  }
  column "username" {
    type = varchar(255)
  }
  column "display_name" {
    type = text
  }
  column "profile_image" {
    type = text
  }
  column "roles" {
    type = sql("role[]")
  }
  column "created_at" {
    type = timestamp
    null = false
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    type = timestamp
    null = false
    default = sql("CURRENT_TIMESTAMP")
  }
  column "deleted_at" {
    type = timestamp
    null = true
  }
  primary_key  {
    columns = [ column.id ]
  }
  unique "username" {
    columns = [ column.username, column.deleted_at ]
  }
}

table "user_passwords" {
  schema = schema.public
  column "user_id" {
    type = uuid 
  }
  column "password" {
    type = varchar(80)
  }
  primary_key {
    columns = [ column.user_id ]
  }
  foreign_key "fk_user_id" {
    columns = [ column.user_id ]
    ref_columns = [ table.users.column.id ]
    on_delete = CASCADE
  }
}

table "courses" {
  schema = schema.public
  column "id" {
    type = uuid
    default = sql("gen_random_uuid()")
  }
  column "name" {
    type = text
  }
  column "created_by" {
    type = uuid
  }
  column "created_at" {
    type = timestamp
  }
  column "updated_at" {
    type = timestamp
  }
  column "deleted_at" {
    type = timestamp
    null = true
  }
  primary_key  {
    columns = [ column.id ]
  }
  foreign_key "fk_created_by" {
    columns = [ column.created_by ]
    ref_columns = [ table.users.column.id ]
  }
  unique "course_name" {
    columns = [ column.name ]
  }
}

table "sections" {
  schema = schema.public
  column "id" {
    type = uuid
    default = sql("gen_random_uuid()")
  }
  column "name" {
    type = text
  }
  column "started_at" {
    type = timestamp
  }
  column "ended_at" {
    type = timestamp
  }
  column "icon" {
    type = text
  }
  column "course_id" {
    type = uuid
  }
  column "semester_id" {
    type = uuid
  }
  primary_key  {
    columns = [ column.id ]
  }
  foreign_key "fk_course_id" {
    columns = [ column.course_id ]
    ref_columns = [ table.courses.column.id ]
  }
  foreign_key "fk_semester_id" {
    columns = [ column.semester_id  ]
    ref_columns = [ table.semesters.column.id ]
  }
    unique "name_course_id" {
    columns = [ column.name,column.course_id ]
  }
}

table "section_instructors" {
  schema = schema.public
  column "section_id" {
    type = uuid
  }
  column "instructor_id" {
    type = uuid
  }
  primary_key  {
    columns = [ column.section_id,  column.instructor_id ]
  }
  foreign_key "fk_section_id" {
    columns = [ column.section_id ]
    ref_columns = [ table.sections.column.id ]
  }
  foreign_key "fk_instructor_id" {
    columns = [ column.instructor_id ]
    ref_columns = [ table.users.column.id ]
  }
}

table "section_tas" {
  schema = schema.public
  column "section_id" {
    type = uuid
  }
  column "ta_id" {
    type = uuid
  }
  primary_key  {
    columns = [ column.section_id,  column.ta_id ]
  }
  foreign_key "fk_section_id" {
    columns = [ column.section_id ]
    ref_columns = [ table.sections.column.id ]
  }
  foreign_key "fk_ta_id" {
    columns = [ column.ta_id ]
    ref_columns = [ table.users.column.id ]
  }
}

table "semesters" {
  schema = schema.public
  column "id" {
    type = uuid
    default = sql("gen_random_uuid()")
  }
  column "name" {
    type = varchar(255)
  }
  column "started_date" {
    type = timestamp
  }
  primary_key  {
    columns = [ column.id ]
  }
  unique "semester_name" {
    columns = [ column.name ]
  }
}

enum "action" {
  schema = schema.public
  values = [ "sign-in" , "sign-out" , "sign-in-failed" ]
}

table "auth_logs" {
  schema = schema.public
  column "id" {
    type = uuid
    default = sql("gen_random_uuid()")
  }
  column "user_id" {
    type = uuid
  }
  column "action" {
    type = enum.action
  }
  column "created_at" {
    type = timestamp
  }
  primary_key  {
    columns = [ column.id ]
  }
  foreign_key "fk_user_id" {
    columns = [ column.user_id ]
    ref_columns = [ table.users.column.id ]
  }
}
