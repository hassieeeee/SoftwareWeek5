#include<stdio.h>
int main(){
    int field[8][8]={
	{0,0,0,0,0,0,0,0},
	{0,0,0,0,0,0,0,0},
	{0,0,0,0,0,0,0,0},
	{0,0,0,1,2,0,0,0},
	{0,0,0,2,1,0,0,0},
	{0,0,0,0,0,0,0,0},
	{0,0,0,0,0,0,0,0},
	{0,0,0,0,0,0,0,0}
    };
    int x,y,n=0,i,j,m,h,k,pn;
    
    printf("  1 2 3 4 5 6 7 8 x\n");
    for(i=0;i<8;i++){
	printf("%d|",i+1);
	for(j=0;j<8;j++){
	    if(field[i][j]==0){
		printf(" |");
	    }else if(field[i][j]==1){
		printf("○|");
	    }else{
		printf("●|");}
	}
	printf("\n");
    }
    printf("y\n");
    printf("オセロです。x座標、y座標の値を入力してバトルしてね。\n");
    printf("1pは○、2pは●です\n");
    printf("x=0,y=0で次の人にパスできます。\n");
    printf("x=11,y=11でゲームを終了できます。\n");
    
for(;;n++){
    	
	pn=n%2+1;
	
	for(;;){
	    printf("player %d\n",pn);
	    printf("x = ");
	    scanf("%d",&y);
	    printf("y = ");
	    scanf("%d",&x);
	    x--;
	    y--;
	    if(0<=x && x<=7 && 0<=y && y<=7 && field[x][y]==0){
		field[x][y]=pn;
	        break;}
	    else if(x== -1 && y== -1 ){
		pn=pn%2+1;}
	    else if(x==10 && y==10){
		break;}
	}
	m=1;/*相手の駒をひっくり返したかどうか判定*/
	if(x== 10 && y== 10){
	    break;}

/*上*/
	for(j=y-1;j>=0;j--){
	    if(field[x][j]==0){
		break;}
	    else if(field[x][j]==pn){
		if(y>j+1){
		    for(h=y-1;h>j;h--){
			field[x][h]=pn;
		    }
		    m=0;
		}
	        break;
	    }
	}
/*下*/
	for(j=y+1;j<=7;j++){
	    if(field[x][j]==0){
		break;}
	    else if(field[x][j]==pn){
		if(y<j-1){
		    for(h=y+1;h<j;h++){
			field[x][h]=pn;
		    }
		    m=0;
		}
	        break;
	    }
	}
/*左*/
	for(i=x-1;i>=0;i--){
	    if(field[i][y]==0){
		break;}
	    else if(field[i][y]==pn){
		if(x>i+1){
		    for(h=x-1;h>i;h--){
			field[h][y]=pn;
		    }
		    m=0;
		}
	        break;
	    }
	}
/*右*/
	for(i=x+1;i<=7;i++){
	    if(field[i][y]==0){
		break;}
	    else if(field[i][y]==pn){
		if(x<i-1){
		    for(h=x+1;h<i;h++){
			field[h][y]=pn;
		    }
		    m=0;
		}
	        break;
	    }
	}
/*右下*/
	for(i=x+1,j=y+1;i<=7&&j<=7;i++,j++){
	    if(field[i][j]==0){
		break;}
	    else if(field[i][j]==pn){
		if(x<i-1){
		    for(h=x+1,k=y+1;h<i;h++,k++){
			field[h][k]=pn;
		    }
		    m=0;
		}
	        break;
	    }
	}
/*左上*/
	for(i=x-1,j=y-1;i>=0&&j>=0;i--,j--){
	    if(field[i][j]==0){
		break;}
	    else if(field[i][j]==pn){
		if(x>i+1){
		    for(h=x-1,k=y-1;h>i;h--,k--){
			field[h][k]=pn;
		    }
		    m=0;
		}
	        break;
	    }
	}
/*右上*/
	for(i=x+1,j=y-1;i<=7&&j>=0;i++,j--){
	    if(field[i][j]==0){
		break;}
	    else if(field[i][j]==pn){
		if(x<i-1){
		    for(h=x+1,k=y-1;h<i;h++,k--){
			field[h][k]=pn;
		    }
		    m=0;
		}
	    break;
	    }
	}
/*左下*/
	for(i=x-1,j=y+1;i>=0&&j<=7;i--,j++){
	    if(field[i][j]==0){
		break;}
	    else if(field[i][j]==pn){
		if(x>i+1){
		    for(h=x-1,k=y+1;h>i;h--,k++){
			field[h][k]=pn;
		    }
		    m=0;
		}
	        break;
	    }
	}
/*不正摘発*/
	if(m==1){
	    field[x][y]=0;
	    n--;
	    printf("不正な位置です。もう一度入力してください。\n");
	}
/*表示*/
	printf(" 1 2 3 4 5 6 7 8 x\n");
	for(i=0;i<8;i++){
	    printf("%d|",i+1);
	    for(j=0;j<8;j++){
		if(field[i][j]==0){
		    printf(" |");
		}else if(field[i][j]==1){
		    printf("○|");
		}else{
		    printf("●|");}
	    }
	    printf("\n");
	}
	printf("y\n");
/*決着判定*/
	int c=0;
	for(i=0;i<=7;i++){
	    for(j=0;j<=7;j++){
		if(field[i][j]==0){
		    c++;}
	    }
	}
	if(c==0){
	    break;}
}
    int pa=0,pb=0;
    for(i=0;i<=7;i++){
	for(j=0;j<=7;j++){
	    if(field[i][j]==1){
		pa++;}
	    else if(field[i][j]==2){
		pb++;}
	}
    }
    n=0;
    printf("%d 対 %d\n",pa,pb);
    if(pa>pb){
	printf("player 1 win!!\n");}
    if(pa<pb){
	printf("player 2 win!!\n");}
    if(pa==pb){
	printf("draw!!\n");}
}