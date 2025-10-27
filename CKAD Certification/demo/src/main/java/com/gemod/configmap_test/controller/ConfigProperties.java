package com.gemod.configmap_test.controller;

import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Component;

@Component
@ConfigurationProperties("greeting")
@Data
public class ConfigProperties {

    private String message;
    private String name;
}
