package usecases

import (
	"reflect"
	"testing"

	"github.com/sugar-cat7/vspo-common-api/domain/entities"
	"github.com/sugar-cat7/vspo-common-api/domain/services"
	"github.com/sugar-cat7/vspo-common-api/usecases/mappers"
)

func TestNewGetClipsByPeriod(t *testing.T) {
	type args struct {
		clipService services.ClipService
		clipMapper  *mappers.ClipMapper
	}
	tests := []struct {
		name string
		args args
		want *GetClipsByPeriod
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGetClipsByPeriod(tt.args.clipService, tt.args.clipMapper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGetClipsByPeriod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetClipsByPeriod_Execute(t *testing.T) {
	type fields struct {
		clipService services.ClipService
		clipMapper  *mappers.ClipMapper
	}
	type args struct {
		start string
		end   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entities.Video
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GetClipsByPeriod{
				clipService: tt.fields.clipService,
				clipMapper:  tt.fields.clipMapper,
			}
			got, err := g.Execute(tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClipsByPeriod.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClipsByPeriod.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
