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

# 修复历史bug
alter table global_settings
    modify pat_text varchar(255) default '' null comment '拍一拍文本';
alter table global_settings
    add pat_voice_timbre varchar(100) default '' null comment '拍一拍音色' after pat_text;