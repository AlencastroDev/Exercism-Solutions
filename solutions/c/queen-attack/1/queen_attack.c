#include "queen_attack.h"


    attack_status_t can_attack(position_t queen_1, position_t queen_2) {
        
        if ( queen_1.row == queen_2.row && queen_1.column == queen_2.column ) {
            return INVALID_POSITION;
        }  
            
        if ( queen_1.row > 7 || queen_2.row > 7 || queen_1.column > 7 || queen_2.column > 7) { 
            return INVALID_POSITION;
        }

        if ( queen_1.column == queen_2.column) {
            return CAN_ATTACK;
        } 

        if ( queen_1.row == queen_2.row) {
            return CAN_ATTACK;
        } 

        
        
        int i = queen_1.row;
        int j = queen_1.column;
            
        while (i > 0 && j < 7) {
            i--;
            j++;

            if (queen_2.row == i && queen_2.column == j) 
                return CAN_ATTACK;
        }
        
        i = queen_1.row;
        j = queen_1.column;
        
        while (i < 7 && j < 7) {
            i++;
            j++;

            if (queen_2.row == i && queen_2.column == j) 
                return CAN_ATTACK;
        }

        i = queen_1.row;
        j = queen_1.column;
        
        while (i < 7 && j > 0) {
            i++;
            j--;

            if (queen_2.row == i && queen_2.column == j) 
                return CAN_ATTACK;
        }

        i = queen_1.row;
        j = queen_1.column;
        
        while (i > 0 && j > 0) {
            i--;
            j--;

            if (queen_2.row == i && queen_2.column == j) 
                return CAN_ATTACK;
        }
        
        return CAN_NOT_ATTACK;
        
     }
       








