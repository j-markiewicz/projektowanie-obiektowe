program bubblesort;

var
	numbers: array of integer;

procedure generate(min, max, count: integer);
var
	i: integer;
begin
	SetLength(numbers, count);

	for i := 0 to High(numbers) do begin
		numbers[i] := random(max + 1 - min) + min;
	end;
end;

procedure sort();
var
	i: integer;
	j: integer;
	temp: integer;
begin
	for i := 0 to High(numbers) do begin
		for j := 1 to High(numbers) - i do begin
			if numbers[j - 1] > numbers[j] then begin
				temp := numbers[j - 1];
				numbers[j - 1] := numbers[j];
				numbers[j] := temp;
			end;
		end;
	end;
end;

procedure testGenMin();
var
	i: integer;
	min: integer;
begin
	for min := 0 to 100 do begin
		generate(min, 100, 100);

		for i := 0 to High(numbers) do begin
			if not (numbers[i] >= min) then begin
				WriteLn('testGenMin failed, found number ', numbers[i], ' with min = ', min);
				Halt(1);
			end;
		end;
	end;
end;

procedure testGenMax();
var
	i: integer;
	max: integer;
begin
	for max := 0 to 100 do begin
		generate(0, max, 100);

		for i := 0 to High(numbers) do begin
			if not (numbers[i] <= max) then begin
				WriteLn('testGenMax failed, found number ', numbers[i], ' with max = ', max);
				Halt(1);
			end;
		end;
	end;
end;

procedure testGenCount();
var
	count: integer;
begin
	for count := 0 to 100 do begin
		generate(0, 100, count);

		if not (Length(numbers) = count) then begin
			WriteLn('testGenCount failed, length ', Length(numbers), ' does not match expected length ', count);
			Halt(1);
		end;
	end;
end;

procedure testSort();
var
	i: integer;
	last: integer;
begin
	SetLength(numbers, 10);

	numbers[0] := 1;
	numbers[1] := 4;
	numbers[2] := 7;
	numbers[3] := 2;
	numbers[4] := 5;
	numbers[5] := 8;
	numbers[6] := 3;
	numbers[7] := 6;
	numbers[8] := 9;
	numbers[9] := 0;

	sort();

	last := -1;
	for i := 0 to 10 do begin
		if not (last < numbers[i]) then begin
			WriteLn('testSort failed for array of 10 numbers');
			Halt(1);
		end;
	end;
end;

procedure testSortEdgecases();
begin
	SetLength(numbers, 0);

	sort();
	if not (Length(numbers) = 0) then begin
		WriteLn('testSortEdgecases failed for 0-length array');
		Halt(1);
	end;

	SetLength(numbers, 1);
	numbers[0] := 123;

	sort();
	if not (Length(numbers) = 1) then begin
		WriteLn('testSortEdgecases failed for 1-length array');
		Halt(1);
	end;
	if not (numbers[0] = 123) then begin
		WriteLn('testSortEdgecases failed for 1-length array');
		Halt(1);
	end;

	SetLength(numbers, 2);
	numbers[0] := 123;
	numbers[1] := 123;

	sort();
	if not (Length(numbers) = 2) then begin
		WriteLn('testSortEdgecases failed for 2-length array of equal values');
		Halt(1);
	end;
	if not (numbers[0] = 123) then begin
		WriteLn('testSortEdgecases failed for 2-length array of equal values');
		Halt(1);
	end;
	if not (numbers[1] = 123) then begin
		WriteLn('testSortEdgecases failed for 2-length array of equal values');
		Halt(1);
	end;
end;

begin
	testGenMin();
	testGenMax();
	testGenCount();
	testSort();
	testSortEdgecases();
	WriteLn('All tests passed');
end.
