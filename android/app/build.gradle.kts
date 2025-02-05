plugins {
    id("org.openapi.generator") version "7.9.0"
    kotlin("plugin.serialization") version "1.9.23"
    alias(libs.plugins.android.application)
    alias(libs.plugins.kotlin.android)
}

val openApiGeneratedDir = "$buildDir/generated/sources/openApi"
openApiGenerate {
    generatorName.set("kotlin")
    library.set("jvm-okhttp4")
   // inputSpec.set("/home/maniu/atm/splitbit/openapi.yml")
    inputSpec.set("/home/golubmik/Documents/ISDC/splitbit/openapi.yml")
    outputDir.set(openApiGeneratedDir)
    packageName.set("fi.isdc_helsinki.splitbit.client")
    // ignoreFileOverride.set(".openapi-generator-ignore")
    configOptions.set(mapOf(
        "serializationLibrary" to "kotlinx_serialization"
    ))
}

android {
    namespace = "fi.isdc_helsinki.splitbit"
    compileSdk = 34

    defaultConfig {
        applicationId = "fi.isdc_helsinki.splitbit"
        minSdk = 29
        targetSdk = 34
        versionCode = 1
        versionName = "1.0"

        testInstrumentationRunner = "androidx.test.runner.AndroidJUnitRunner"
        vectorDrawables {
            useSupportLibrary = true
        }
    }
    sourceSets {
        getByName("main") {
            kotlin.srcDir("${openApiGeneratedDir}/src/main/kotlin")
        }
    }
    buildTypes {
        release {
            isMinifyEnabled = true
            proguardFiles(
                getDefaultProguardFile("proguard-android-optimize.txt"),
                "proguard-rules.pro"
            )
            signingConfig = signingConfigs.getByName("debug")
            ndk {
                abiFilters.add("arm64-v8a")
            }
        }
    }
    compileOptions {
        sourceCompatibility = JavaVersion.VERSION_1_8
        targetCompatibility = JavaVersion.VERSION_1_8
    }
    kotlinOptions {
        jvmTarget = "1.8"
    }
    buildFeatures {
        compose = true
    }
    composeOptions {
        kotlinCompilerExtensionVersion = "1.5.1"
    }
    packaging {
        resources {
            excludes += "/META-INF/{AL2.0,LGPL2.1}"
        }
    }
}

dependencies {
    // These are the dependencies for the autogenerated models
    implementation("com.squareup.okhttp3:okhttp:4.12.0")
    implementation("org.jetbrains.kotlinx:kotlinx-serialization-json:1.6.2")

    implementation("com.squareup.retrofit2:converter-kotlinx-serialization:2.11.0")
    implementation("com.squareup.retrofit2:converter-scalars:2.11.0")
    implementation("com.squareup.okhttp3:logging-interceptor:4.12.0")
    implementation("com.squareup.retrofit2:retrofit:2.11.0")
    implementation("com.squareup.retrofit2:converter-gson:2.11.0")
    implementation(libs.androidx.core.ktx)
    implementation(libs.androidx.lifecycle.runtime.ktx)
    implementation(libs.androidx.activity.compose)
    implementation(platform(libs.androidx.compose.bom))
    implementation(libs.androidx.ui)
    implementation(libs.androidx.ui.graphics)
    implementation(libs.androidx.ui.tooling.preview)
    implementation(libs.androidx.material3)
    implementation(libs.androidx.compose.material.iconsExtended)
    implementation(libs.androidx.navigation.runtime.ktx)
    implementation(libs.androidx.navigation.compose)
    implementation("androidx.security:security-crypto:1.1.0-alpha06")
    implementation("com.auth0.android:jwtdecode:2.0.1")
    testImplementation(libs.junit)
    androidTestImplementation(libs.androidx.junit)
    androidTestImplementation(libs.androidx.espresso.core)
    androidTestImplementation(platform(libs.androidx.compose.bom))
    androidTestImplementation(libs.androidx.ui.test.junit4)
    debugImplementation(libs.androidx.ui.tooling)
    debugImplementation(libs.androidx.ui.test.manifest)

}
