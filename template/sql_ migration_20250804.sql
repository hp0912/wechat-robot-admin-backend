alter table global_settings
    add workflow_model varchar(100) default '' null comment '工作流AI模型名称' after chat_api_key;

alter table global_settings
    add image_recognition_model varchar(100) default '' null comment '图像识别模型名称' after chat_model;

alter table friend_settings
    add workflow_model varchar(100) default '' null comment '工作流AI模型名称' after chat_api_key;

alter table friend_settings
    add image_recognition_model varchar(100) default '' null comment '图像识别模型名称' after chat_model;

alter table chat_room_settings
    add workflow_model varchar(100) default '' null comment '工作流AI模型名称' after chat_api_key;

alter table chat_room_settings
    add image_recognition_model varchar(100) default '' null comment '图像识别模型名称' after chat_model;