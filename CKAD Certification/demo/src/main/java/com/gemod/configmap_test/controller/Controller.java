package com.gemod.configmap_test.controller;


import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api")
public class Controller {

    private ConfigProperties configProperties;

    public Controller(ConfigProperties configProperties) {
        this.configProperties = configProperties;
    }

    @GetMapping
    public ResponseEntity<String> get() {
        String envVar = System.getenv().getOrDefault("ENV_VAR", "Hi");
        String configmapValues = String.format(" %s, it's %s here!", configProperties.getMessage(), configProperties.getName());
        StringBuilder result = new StringBuilder(configmapValues).append("\n").append(String.format("The value of env var is \"%s\"", envVar));
        return ResponseEntity.ok(result.toString());
    }
}
