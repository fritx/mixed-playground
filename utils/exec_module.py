import importlib.util, os, inspect
from types import ModuleType

def exec_module(path: str) -> ModuleType:

    call_file = get_previous_frame_file(2)

    # fix: FileNotFoundError: [Errno 2] No such file or directory: 'xxx/mixed-playground/xxx.py'
    current_file = os.path.abspath(path=call_file)
    current_dir = os.path.dirname(current_file)
    src_file = os.path.join(current_dir, path)

    spec = importlib.util.spec_from_file_location("x", src_file)
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)
    return module

# @author 文心一言
def get_previous_frame_file(offset: int = 1) -> str:

    # 获取当前帧的堆栈信息
    current_frame = inspect.currentframe()
    previous_frame = current_frame

    # 获取上一层的帧对象
    while offset:
        offset -= 1
        previous_frame = previous_frame.f_back

    # 从上一层的帧对象中获取文件名
    previous_file = previous_frame.f_code.co_filename

    return previous_file
