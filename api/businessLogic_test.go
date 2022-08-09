package api

import (
	"reflect"
	"testing"
)

func Test_calculateBusinessLogic(t *testing.T) {
	type args struct {
		userOrderCount            int8
		userItemCount             int
		userBasketTotal           float32
		basketTotalGIVEN          float32
		userInMonthTotal          float32
		totalPurchaseInMonthGIVEN float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		// Test cases for busines logic.
		{
			args: args{
				userOrderCount:            3,
				userItemCount:             1,
				userBasketTotal:           1,
				basketTotalGIVEN:          0,
				userInMonthTotal:          0,
				totalPurchaseInMonthGIVEN: 1,
			},
			want: []float32{0, .9, .85},
		},
		{
			args: args{
				userOrderCount:            3,
				userItemCount:             1,
				userBasketTotal:           1,
				basketTotalGIVEN:          0,
				userInMonthTotal:          1,
				totalPurchaseInMonthGIVEN: 0,
			},
			want: []float32{.9, .9, .85},
		},
		{
			args: args{
				userOrderCount:            3,
				userItemCount:             4,
				userBasketTotal:           1,
				basketTotalGIVEN:          0,
				userInMonthTotal:          1,
				totalPurchaseInMonthGIVEN: 0,
			},
			want: []float32{.92, .9, .85},
		},
		{
			args: args{
				userOrderCount:            1,
				userItemCount:             1,
				userBasketTotal:           1,
				basketTotalGIVEN:          0,
				userInMonthTotal:          1,
				totalPurchaseInMonthGIVEN: 0,
			},
			want: []float32{.9, .9, .9},
		},
		{
			args: args{
				userOrderCount:            1,
				userItemCount:             4,
				userBasketTotal:           1,
				basketTotalGIVEN:          0,
				userInMonthTotal:          0,
				totalPurchaseInMonthGIVEN: 1,
			},
			want: []float32{.92, .92, .92},
		},
		{
			args: args{
				userOrderCount:            1,
				userItemCount:             1,
				userBasketTotal:           0,
				basketTotalGIVEN:          1,
				userInMonthTotal:          0,
				totalPurchaseInMonthGIVEN: 1,
			},
			want: []float32{0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateBusinessLogic(tt.args.userOrderCount, tt.args.userItemCount, tt.args.userBasketTotal, tt.args.basketTotalGIVEN, tt.args.userInMonthTotal, tt.args.totalPurchaseInMonthGIVEN); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateBusinessLogic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateDiscount(t *testing.T) {
	type args struct {
		b        []float32
		c        []int16
		discount []float32
	}
	tests := []struct {
		name               string
		args               args
		wantCalculatePrice float32
		wantCalculateVAT   float32
	}{
		{
			args: args{
				b:        []float32{2200, 35.75, 14.52},
				c:        []int16{18, 8, 1},
				discount: []float32{389.40, 3.861, 0},
			},
			wantCalculatePrice: 5792.1187,
			wantCalculateVAT:   67.58348,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCalculatePrice, gotCalculateVAT := calculateDiscount(tt.args.b, tt.args.c, tt.args.discount)
			if gotCalculatePrice != tt.wantCalculatePrice {
				t.Errorf("calculateDiscount() gotCalculatePrice = %v, want %v", gotCalculatePrice, tt.wantCalculatePrice)
			}
			if gotCalculateVAT != tt.wantCalculateVAT {
				t.Errorf("calculateDiscount() gotCalculateVAT = %v, want %v", gotCalculateVAT, tt.wantCalculateVAT)
			}
		})
	}
}
