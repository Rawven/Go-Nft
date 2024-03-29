package com.topview.util;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.extern.slf4j.Slf4j;

import java.util.List;
import java.util.Map;

/**
 * json util
 *
 * @author 刘家辉
 * @date 2023/11/23
 */
@Slf4j
public class JsonUtil {
    public static final ObjectMapper OBJECT_MAPPER;

    static {
        OBJECT_MAPPER = new ObjectMapper();
    }

    public static List<Object> strToList(String str) {
        try {
            return OBJECT_MAPPER.readValue(str, new TypeReference<>() {
            });
        } catch (JsonProcessingException e) {
            log.error("JsonUtil strToObj error:{}", e.getMessage());
            throw new RuntimeException(e);
        }
    }
    public static Map jsonToMap(String json) {
        try {
            return OBJECT_MAPPER.readValue(json, Map.class);
        } catch (JsonProcessingException e) {
            log.error("JsonUtil jsonToMap error:{}", e.getMessage());
            throw new RuntimeException(e);
        }
    }

    public static <T> T jsonToObj(String json, Class<T> clazz) {
        try {
            return OBJECT_MAPPER.readValue(json, clazz);
        } catch (JsonProcessingException e) {
            log.error("JsonUtil jsonToObj error:{}", e.getMessage());
            throw new RuntimeException(e);
        }
    }

    public static <T> String objToJson(T obj) {
        try {
            return OBJECT_MAPPER.writeValueAsString(obj);
        } catch (JsonProcessingException e) {
            log.error("JsonUtil objToJson error:{}", e.getMessage());
            throw new RuntimeException(e);
        }

    }
}
