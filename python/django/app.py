import os
import sys
from django.conf import settings
from django.conf.urls import patterns, include, url
from django.http import HttpResponse
from django.core.management import execute_from_command_line
 
filename = os.path.splitext(os.path.basename(__file__))[0]

urlpatterns = patterns(
    '',
    url(r'psql/select', '%s.psql_select' % filename, name='home'),
)


def psql_select(request):
    return HttpResponse('Django rules!')
 

if __name__ == "__main__":
    os.environ.setdefault("DJANGO_SETTINGS_MODULE", "settings")
    execute_from_command_line([sys.argv[0], 'runserver'])
