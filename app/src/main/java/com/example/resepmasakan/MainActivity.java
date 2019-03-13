package com.example.resepmasakan;

import android.content.Intent;
import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.widget.Button;


public class MainActivity extends AppCompatActivity implements View.OnClickListener{
    private Button btnMoveActivity;
    private Button btnMoveActivity2;
    private Button btnMoveActivity3;
    private Button btnMoveActivity4;
    private Button btnMoveActivity5;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        btnMoveActivity = (Button) findViewById(R.id.recipe_1_btn);
        btnMoveActivity.setOnClickListener(this);
        btnMoveActivity2 = (Button) findViewById(R.id.recipe_2_btn);
        btnMoveActivity2.setOnClickListener(this);
        btnMoveActivity3 = (Button) findViewById(R.id.recipe_3_btn);
        btnMoveActivity3.setOnClickListener(this);
        btnMoveActivity4 = (Button) findViewById(R.id.recipe_4_btn);
        btnMoveActivity4.setOnClickListener(this);
        btnMoveActivity5 = (Button) findViewById(R.id.recipe_4_btn);
        btnMoveActivity5.setOnClickListener(this);

    }

    @Override
    public void onClick(View view){
        switch (view.getId()){
            case R.id.recipe_1_btn:
                Intent moveIntent = new Intent(MainActivity.this, OsengCumiKemangiActivity.class);
                startActivity(moveIntent);
                break;
            case R.id.recipe_2_btn:
                Intent moveIntent2 = new Intent(MainActivity.this, SopIgaSapiActivity.class);
                startActivity(moveIntent2);
                break;
            case R.id.recipe_3_btn:
                Intent moveIntent3 = new Intent(MainActivity.this, SopIgaSapiActivity.class);
                startActivity(moveIntent3);
                break;
            case R.id.recipe_4_btn:
                Intent moveIntent4 = new Intent(MainActivity.this, MieGorengPedasActivity.class);
                startActivity(moveIntent4);
                break;
            case R.id.recipe_5_btn:
                Intent moveIntent5 = new Intent(MainActivity.this, UdangSambalIrisActivity.class);
                startActivity(moveIntent5);
                }

        }
    }
