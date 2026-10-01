"""Exercise CLI failures and extraction with POSIX resource limits."""
import pathlib
import resource
import signal
import subprocess
import tempfile
import zipfile


def limit_files():
    resource.setrlimit(resource.RLIMIT_NOFILE, (64, 64))


def limit_size():
    resource.setrlimit(resource.RLIMIT_FSIZE, (100, 100))
    signal.signal(signal.SIGXFSZ, signal.SIG_IGN)


with tempfile.TemporaryDirectory() as directory:
    root = pathlib.Path(directory)
    binary = root / "zipl"
    subprocess.run(["go", "build", "-o", str(binary), "./cmd/zipl"], check=True)
    source = root / "input"
    source.mkdir()
    (source / "empty").mkdir()
    for index in range(100):
        (source / f"{index}.txt").write_text("hello" * 100)
    for spelling in ("input", "input/", "./input"):
        output = root / "out.zip"
        result = subprocess.run(
            [str(binary), "-t", spelling, "-o", str(output)],
            cwd=root, preexec_fn=limit_files, capture_output=True, text=True,
        )
        assert result.returncode == 0, result.stderr
        with zipfile.ZipFile(output) as archive:
            assert len(archive.infolist()) == 102
            assert archive.getinfo("input/").is_dir()
            assert archive.getinfo("input/empty/").is_dir()
            archive.extractall(root / "extracted")
        assert (root / "extracted/input/0.txt").read_text() == "hello" * 100
        assert (root / "extracted/input/empty").is_dir()
    result = subprocess.run(
        [str(binary), "-t", str(source / "0.txt"), "-o", str(root / "limited.zip")],
        preexec_fn=limit_size, capture_output=True, text=True,
    )
    assert result.returncode != 0, "CLI silently accepted a truncated ZIP"
    assert result.stderr, "CLI should report the final write failure"
    assert (root / "limited.zip").stat().st_size == 100
print("CLI resource limits and Python ZIP extraction passed")
