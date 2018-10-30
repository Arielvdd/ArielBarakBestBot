// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testCommandsChannelsOverrides(t *testing.T) {
	t.Parallel()

	query := CommandsChannelsOverrides()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCommandsChannelsOverridesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CommandsChannelsOverride{}
	if err = randomize.Struct(seed, o, commandsChannelsOverrideDBTypes, true, commandsChannelsOverrideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CommandsChannelsOverrides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCommandsChannelsOverridesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CommandsChannelsOverride{}
	if err = randomize.Struct(seed, o, commandsChannelsOverrideDBTypes, true, commandsChannelsOverrideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := CommandsChannelsOverrides().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CommandsChannelsOverrides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCommandsChannelsOverridesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CommandsChannelsOverride{}
	if err = randomize.Struct(seed, o, commandsChannelsOverrideDBTypes, true, commandsChannelsOverrideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CommandsChannelsOverrideSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CommandsChannelsOverrides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCommandsChannelsOverridesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CommandsChannelsOverride{}
	if err = randomize.Struct(seed, o, commandsChannelsOverrideDBTypes, true, commandsChannelsOverrideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CommandsChannelsOverrideExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if CommandsChannelsOverride exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CommandsChannelsOverrideExists to return true, but got false.")
	}
}

func testCommandsChannelsOverridesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CommandsChannelsOverride{}
	if err = randomize.Struct(seed, o, commandsChannelsOverrideDBTypes, true, commandsChannelsOverrideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	commandsChannelsOverrideFound, err := FindCommandsChannelsOverride(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if commandsChannelsOverrideFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCommandsChannelsOverridesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CommandsChannelsOverride{}
	if err = randomize.Struct(seed, o, commandsChannelsOverrideDBTypes, true, commandsChannelsOverrideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = CommandsChannelsOverrides().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCommandsChannelsOverridesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CommandsChannelsOverride{}
	if err = randomize.Struct(seed, o, commandsChannelsOverrideDBTypes, true, commandsChannelsOverrideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := CommandsChannelsOverrides().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCommandsChannelsOverridesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	commandsChannelsOverrideOne := &CommandsChannelsOverride{}
	commandsChannelsOverrideTwo := &CommandsChannelsOverride{}
	if err = randomize.Struct(seed, commandsChannelsOverrideOne, commandsChannelsOverrideDBTypes, false, commandsChannelsOverrideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}
	if err = randomize.Struct(seed, commandsChannelsOverrideTwo, commandsChannelsOverrideDBTypes, false, commandsChannelsOverrideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = commandsChannelsOverrideOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = commandsChannelsOverrideTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CommandsChannelsOverrides().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCommandsChannelsOverridesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	commandsChannelsOverrideOne := &CommandsChannelsOverride{}
	commandsChannelsOverrideTwo := &CommandsChannelsOverride{}
	if err = randomize.Struct(seed, commandsChannelsOverrideOne, commandsChannelsOverrideDBTypes, false, commandsChannelsOverrideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}
	if err = randomize.Struct(seed, commandsChannelsOverrideTwo, commandsChannelsOverrideDBTypes, false, commandsChannelsOverrideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = commandsChannelsOverrideOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = commandsChannelsOverrideTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CommandsChannelsOverrides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testCommandsChannelsOverridesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CommandsChannelsOverride{}
	if err = randomize.Struct(seed, o, commandsChannelsOverrideDBTypes, true, commandsChannelsOverrideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CommandsChannelsOverrides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCommandsChannelsOverridesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CommandsChannelsOverride{}
	if err = randomize.Struct(seed, o, commandsChannelsOverrideDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(commandsChannelsOverrideColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := CommandsChannelsOverrides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCommandsChannelsOverrideToManyCommandsCommandOverrides(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CommandsChannelsOverride
	var b, c CommandsCommandOverride

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, commandsChannelsOverrideDBTypes, true, commandsChannelsOverrideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, commandsCommandOverrideDBTypes, false, commandsCommandOverrideColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, commandsCommandOverrideDBTypes, false, commandsCommandOverrideColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.CommandsChannelsOverridesID = a.ID
	c.CommandsChannelsOverridesID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	commandsCommandOverride, err := a.CommandsCommandOverrides().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range commandsCommandOverride {
		if v.CommandsChannelsOverridesID == b.CommandsChannelsOverridesID {
			bFound = true
		}
		if v.CommandsChannelsOverridesID == c.CommandsChannelsOverridesID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CommandsChannelsOverrideSlice{&a}
	if err = a.L.LoadCommandsCommandOverrides(ctx, tx, false, (*[]*CommandsChannelsOverride)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CommandsCommandOverrides); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.CommandsCommandOverrides = nil
	if err = a.L.LoadCommandsCommandOverrides(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CommandsCommandOverrides); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", commandsCommandOverride)
	}
}

func testCommandsChannelsOverrideToManyAddOpCommandsCommandOverrides(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CommandsChannelsOverride
	var b, c, d, e CommandsCommandOverride

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, commandsChannelsOverrideDBTypes, false, strmangle.SetComplement(commandsChannelsOverridePrimaryKeyColumns, commandsChannelsOverrideColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*CommandsCommandOverride{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, commandsCommandOverrideDBTypes, false, strmangle.SetComplement(commandsCommandOverridePrimaryKeyColumns, commandsCommandOverrideColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*CommandsCommandOverride{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCommandsCommandOverrides(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.CommandsChannelsOverridesID {
			t.Error("foreign key was wrong value", a.ID, first.CommandsChannelsOverridesID)
		}
		if a.ID != second.CommandsChannelsOverridesID {
			t.Error("foreign key was wrong value", a.ID, second.CommandsChannelsOverridesID)
		}

		if first.R.CommandsChannelsOverride != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.CommandsChannelsOverride != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.CommandsCommandOverrides[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.CommandsCommandOverrides[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.CommandsCommandOverrides().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCommandsChannelsOverridesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CommandsChannelsOverride{}
	if err = randomize.Struct(seed, o, commandsChannelsOverrideDBTypes, true, commandsChannelsOverrideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCommandsChannelsOverridesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CommandsChannelsOverride{}
	if err = randomize.Struct(seed, o, commandsChannelsOverrideDBTypes, true, commandsChannelsOverrideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CommandsChannelsOverrideSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCommandsChannelsOverridesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CommandsChannelsOverride{}
	if err = randomize.Struct(seed, o, commandsChannelsOverrideDBTypes, true, commandsChannelsOverrideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CommandsChannelsOverrides().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	commandsChannelsOverrideDBTypes = map[string]string{`AutodeleteResponse`: `boolean`, `AutodeleteResponseDelay`: `integer`, `AutodeleteTrigger`: `boolean`, `AutodeleteTriggerDelay`: `integer`, `ChannelCategories`: `ARRAYbigint`, `Channels`: `ARRAYbigint`, `CommandsEnabled`: `boolean`, `Global`: `boolean`, `GuildID`: `bigint`, `ID`: `bigint`, `IgnoreRoles`: `ARRAYbigint`, `RequireRoles`: `ARRAYbigint`}
	_                               = bytes.MinRead
)

func testCommandsChannelsOverridesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(commandsChannelsOverridePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(commandsChannelsOverrideColumns) == len(commandsChannelsOverridePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CommandsChannelsOverride{}
	if err = randomize.Struct(seed, o, commandsChannelsOverrideDBTypes, true, commandsChannelsOverrideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CommandsChannelsOverrides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, commandsChannelsOverrideDBTypes, true, commandsChannelsOverridePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCommandsChannelsOverridesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(commandsChannelsOverrideColumns) == len(commandsChannelsOverridePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CommandsChannelsOverride{}
	if err = randomize.Struct(seed, o, commandsChannelsOverrideDBTypes, true, commandsChannelsOverrideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CommandsChannelsOverrides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, commandsChannelsOverrideDBTypes, true, commandsChannelsOverridePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(commandsChannelsOverrideColumns, commandsChannelsOverridePrimaryKeyColumns) {
		fields = commandsChannelsOverrideColumns
	} else {
		fields = strmangle.SetComplement(
			commandsChannelsOverrideColumns,
			commandsChannelsOverridePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := CommandsChannelsOverrideSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCommandsChannelsOverridesUpsert(t *testing.T) {
	t.Parallel()

	if len(commandsChannelsOverrideColumns) == len(commandsChannelsOverridePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := CommandsChannelsOverride{}
	if err = randomize.Struct(seed, &o, commandsChannelsOverrideDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CommandsChannelsOverride: %s", err)
	}

	count, err := CommandsChannelsOverrides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, commandsChannelsOverrideDBTypes, false, commandsChannelsOverridePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CommandsChannelsOverride struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CommandsChannelsOverride: %s", err)
	}

	count, err = CommandsChannelsOverrides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
